package optional

import (
	"strconv"
	"math"
	"golang.org/x/net/html"
	"io"
	"log"
	"strings"
	"net/http"
	"fmt"
	"time"
	"database/sql"
    _ "github.com/lib/pq"
	"github.com/spf13/viper"
)
func Gorut() {
	
	for true {
		Full_bd_rate()
		time.Sleep(240 * time.Second) 
	}

}
func Full_bd_rate() {
	var netClient = http.Client{
		Timeout: time.Second * 10,
	}
	    response, err := netClient.Get("https://www.sberometer.ru/cbr/")
		if err != nil {
		log.Fatal(err)
    	}
	defer response.Body.Close()
	
	textTags := []string{
		"a",
	}

	tag := ""
	enter := false
    var numbers [4]float32
    var num int = 0
	tokenizer := html.NewTokenizer(response.Body)
	for {
		tt := tokenizer.Next()
		token := tokenizer.Token()

		err := tokenizer.Err()
		if err == io.EOF {
			break
		}

		switch tt {
		case html.ErrorToken:
			log.Fatal(err)
		case html.StartTagToken, html.SelfClosingTagToken:
			enter = false

			tag = token.Data
			for _, ttt := range textTags {
				if tag == ttt {
					enter = true
					break
				}
			}
		case html.TextToken:
			if enter {
				data := strings.TrimSpace(token.Data)

				if len(data) == 7 {
					if num < 4 {
							value, err := strconv.ParseFloat(data, 32)
					if err != nil {
						fmt.Println(err)
					}
					numbers[num] = float32(value)
					num++
					numbers[num] = float32(math.Ceil((1/float64(value))*100000)/100000)
					num++
					}
				
				}
			}
		}
	}
	if err := initConfig();
	err != nil {
		log.Fatalf(err.Error())
	}
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",viper.GetString("user"),viper.GetString("password"),viper.GetString("dbname"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	} 
	defer db.Close()
	
	today := time.Now()

	update1, err := db.Query(fmt.Sprintf("update curren set rate = '%f', date='%s' where currency1 = 'USD' and currency2 = 'RUB'",numbers[0],today.Format(time.RFC1123)))
	if err != nil {
		panic(err)
	}
	defer update1.Close() 
	
	update2, err := db.Query(fmt.Sprintf("update curren set rate = '%f', date='%s' where currency1 = 'RUB' and currency2 = 'USD'",numbers[1],today.Format(time.RFC1123)))
	if err != nil {
		panic(err)
	}
	defer update2.Close() 
	
	update3, err := db.Query(fmt.Sprintf("update curren set rate = '%f', date='%s' where currency1 = 'EUR' and currency2 = 'RUB'",numbers[2],today.Format(time.RFC1123)))
	if err != nil {
		panic(err)
	}
	defer update3.Close() 

	update4, err := db.Query(fmt.Sprintf("update curren set rate = '%f', date='%s' where currency1 = 'RUB' and currency2 = 'EUR'",numbers[3],today.Format(time.RFC1123)))
	if err != nil {
		panic(err)
	}
	defer update4.Close()

    
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}