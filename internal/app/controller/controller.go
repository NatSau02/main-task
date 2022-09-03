package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/NatSau02/main-task"
	"github.com/NatSau02/main-task/internal/app/service"
	"github.com/NatSau02/main-task/internal/optional"
	"fmt"
	"time"
	"strconv"
	"database/sql"
    _ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"

)

type RateController interface {
	FindAll() []entity.Сalculator
	Save(ctx *gin.Context) entity.Сalculator
}

type controller struct {
	service service.RateServicee
}

func New1(service service.RateServicee) RateController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Сalculator {
	return c.service.FindAll()
}
 


func (c *controller) Save(ctx *gin.Context) entity.Сalculator {
	var rate entity.Сalculator 

    if err := initConfig();
	err != nil {
		log.Fatalf(err.Error())
	}

	ctx.BindJSON(&rate)
	rate.ValueСurrency2 = "2"

  
	//connStr := fmt.Sprintf("user=postgres password=123 dbname=main_project sslmode=disable")
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",viper.GetString("user"),viper.GetString("password"),viper.GetString("dbname"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	} 
	defer db.Close()
	var c1 string = rate.Сurrency1
    var c2 string = rate.Сurrency2
	today := time.Now()
	
	rows, err := db.Query(fmt.Sprintf("select * from curren where currency1 = '%s' and currency2 = '%s'", c1,c2))
    if err != nil {
        panic(err)
    }
    defer rows.Close()
    
    var showPost entity.BD
	for rows.Next(){
        var p entity.BD
        err := rows.Scan(&p.Id,&p.Сurrency1,&p.Сurrency2,&p.Rate,&p.Date)
        if err != nil{
            fmt.Println(err)
            continue
        }
	  showPost = p
	}
	
	
	
	if showPost.Сurrency1 == "" {
		
		insert, err := db.Query(fmt.Sprintf("insert into curren (currency1, currency2, rate, date) values ('%s','%s', '%f','%s')", c1, c2, 75.77, today.Format(time.RFC1123)))
			if err != nil {
				panic(err)
			}
			defer insert.Close() 
	 
		  optional.Full_bd_rate()
		    
	
		rows, err := db.Query(fmt.Sprintf("select * from curren where currency1 = '%s' and currency2 = '%s'", c1,c2))
		if err != nil {
			panic(err)
		}
		defer rows.Close()
     
		var showP entity.BD
		for rows.Next(){
			var p entity.BD
			err := rows.Scan(&p.Id,&p.Сurrency1,&p.Сurrency2,&p.Rate,&p.Date)
			if err != nil{
				fmt.Println(err)
				continue
			}
		showP = p
		 
		}
		if s, err := strconv.ParseFloat(rate.ValueСurrency1, 64); err == nil {
			rate.ValueСurrency2 = fmt.Sprintf("%f", showP.Rate*float32(s))
		}

	} else {
		
		
		if s, err := strconv.ParseFloat(rate.ValueСurrency1, 64); err == nil {
			rate.ValueСurrency2 = fmt.Sprintf("%f", showPost.Rate*float32(s))
		}
	}
	

	
	c.service.Save(rate)
	return rate
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}