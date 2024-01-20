package main

import (
	"fmt"
)

func main(){

	// declare
	dessertPrices := map[string]float64{
        "Chocolate Cake":      15.99,
        "Cheesecake":          12.50,
        "Tiramisu":            18.75,
        "Apple Pie":           14.00,
        "Ice Cream Sundae":    8.99,
        "Creme Brulee":        16.25,
        "Strawberry Shortcake": 13.50,
        "Banana Split":        10.99,
        "Pecan Pie":           17.50,
        "Red Velvet Cupcake":   7.95,
    }

	// --print biasa--
	// fmt.Println(dessertPrices)

	// --loop map--
	// for key, value := range dessertPrices{
	// 	fmt.Printf("%v price is %v$\n", key, value)
	// }

	// --print some index--
	// fmt.Printf("%v$ \n", dessertPrices["Pecan Pie"])
	// harus menginput key didalam map sesuai tipe data "string"

	// --mengubah value pada sebuah key--
	// dessertPrices["Pecan Pie"] = 1.0
	// fmt.Println(dessertPrices)
	// // value Pecan Pie berubah
}