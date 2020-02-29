package main

import (
	"container/list"
	"fmt"
)

type Product struct {
	Name      string `json:"name"`
	ProductID string `json:"product_id"`
	Number    string  `json:"number"`
	Price     string `json:"price"`
	IsOnSale  map[string]string `json:"is_on_sale"`
	s *list.List
}
func main(){
	s := "abc"
	fmt.Println(string(s[1]))
}
