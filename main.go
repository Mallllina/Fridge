package main

import (
	"fmt"
	"fridge/internal/product"
	"fridge/internal/storage"
	"time"
)

func main() {
	fridge := storage.NewFridge(10)
	freezer := storage.NewFreezer(10)

	fridge.Open()
	freezer.Open()

	product1 := product.Product{
		Name:           "Milk",
		Volume:         2.5,
		CreatedAt:      time.Now(),
		ExpirationDate: 7,
	}

	err := fridge.AddProduct(product1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	fridge.Close()

	product2 := product.Product{
		Name:           "Cake",
		Volume:         4,
		CreatedAt:      time.Now(),
		ExpirationDate: 15,
	}

	err = fridge.AddProduct(product2)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	product3 := product.Product{
		Name:           "Meat",
		Volume:         3,
		CreatedAt:      time.Now(),
		ExpirationDate: 6,
	}

	err = freezer.AddProduct(product3)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
