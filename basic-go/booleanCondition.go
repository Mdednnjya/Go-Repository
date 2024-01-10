package main

import (
	"fmt"
	"sort"
)

func main(){
	// declaration
	// grade := 91
	months := []string{
        "January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December",
    }



	//  example-1
	// if grade >= 90 {
	// 	fmt.Println("You recived A grade")	
	// } else if grade >= 80 {
	// 	fmt.Println("You recived B grade")
	// } else if grade >= 70 {
	// 	fmt.Println("You recived C grade")
	// } else if grade >= 60 {
	// 	fmt.Println("You recived D grade")
	// } else if grade >= 50 {
	// 	fmt.Println("You recived E grade")
	// } else {
	// 	fmt.Println("You yet to pass")
	// }

	//  example-2: loop with break example
	fmt.Println("Kita akan mencari bulan October")
	for index, _ := range months{
		searchedMonth := sort.SearchStrings(months, "October")
		fmt.Printf("kita sekarang berada pada bulan ke %v yaitu %v \n", (index + 1), months[index])
		if index == (searchedMonth + 1) {
			fmt.Println("")
			fmt.Printf("Kita sudah mendapatkan bulan October\n")
			fmt.Printf("Loop Berakhir")
			break
		}
		index++
	}
}