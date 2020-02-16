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
	a := [][]int{
		{1,2,3},
		{2,3,3},
		{4,6,5},
	}
	for _ , l := range  a {
		fmt.Println(l[0])
	}
}