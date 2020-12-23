package main

import (
	"testing"
)

//Method to add items to the basket
func addProductToBasket(basket *productList, productCode string) {
	for _, currentProduct := range products {
		if currentProduct.Code == productCode {
			*basket = append(*basket, currentProduct)
		}
	}
}
func Test1(t *testing.T) {
	basket := productList{}
	total := 0.0
	addProductToBasket(&basket, "PEN")
	addProductToBasket(&basket, "TSHIRT")
	addProductToBasket(&basket, "MUG")
	calculateTotal(basket, &total)
	if total != 32.50 {
		t.Errorf("Expected 32.50, got %v", total)
	}
}

func Test2(t *testing.T) {
	basket := productList{}
	total := 0.0
	addProductToBasket(&basket, "PEN")
	addProductToBasket(&basket, "TSHIRT")
	addProductToBasket(&basket, "PEN")
	calculateTotal(basket, &total)
	if total != 25.00 {
		t.Errorf("Expected 25.00, got %v", total)
	}
}

func Test3(t *testing.T) {
	basket := productList{}
	total := 0.0
	addProductToBasket(&basket, "TSHIRT")
	addProductToBasket(&basket, "TSHIRT")
	addProductToBasket(&basket, "TSHIRT")
	addProductToBasket(&basket, "PEN")
	addProductToBasket(&basket, "TSHIRT")
	calculateTotal(basket, &total)
	if total != 65.00 {
		t.Errorf("Expected 65.00, got %v", total)
	}
}

func Test4(t *testing.T) {
	basket := productList{}
	total := 0.0
	addProductToBasket(&basket, "PEN")
	addProductToBasket(&basket, "TSHIRT")
	addProductToBasket(&basket, "PEN")
	addProductToBasket(&basket, "PEN")
	addProductToBasket(&basket, "MUG")
	addProductToBasket(&basket, "TSHIRT")
	addProductToBasket(&basket, "TSHIRT")
	calculateTotal(basket, &total)
	if total != 62.50 {
		t.Errorf("Expected 62.50, got %v", total)
	}
}
