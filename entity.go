package entity

type BD struct {
	Id       int 
	Сurrency1 string 
	Сurrency2 string  
	Rate float32       
	Date string  
}
type Сalculator struct {
	Сurrency1       string `json:"currency1"`
	ValueСurrency1 string `json:"valueСurrency1"`
	Сurrency2          string `json:"currency2"`
	ValueСurrency2    string `json:"valueСurrency2"`
}