package main

import "fmt"

type Product struct {
	Name      string `json:"name"`
	ProductID string `json:"product_id"`
	Number    string  `json:"number"`
	Price     string `json:"price"`
	IsOnSale  map[string]string `json:"is_on_sale"`
}
func main(){
		fmt.Println(2 % 3)
}