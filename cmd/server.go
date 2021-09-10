package main

import (
	"fmt"

	"github.com/halhal23/strategy-product/domain/model"
)

func main(){
	product := model.NewProductModel(1, "PC", 18_000, 2)
	fmt.Println("hello product")
	fmt.Println(product)
}