package main

import (
	"fmt"
)

func main(){
	// declare variable
	// x := 0
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

	// [1] --basic interation--
	// for x < 4 {
	// 	fmt.Println("iterasi ke-", (x + 1))
	// 	x++
	// 	x++
	// 	// 2 times iteration since "x++" written 2 times
	// }

	// [2] --tradisional for_loop--
	// example-1
	// for y := 5; y > 0;y--{
	// 	fmt.Println((0 + y), "loop lagi")
	// }

	// example-2
	// for bulan := 0; bulan < len(months); bulan++{
	// 	fmt.Println((bulan + 1),"", months[bulan])
	// }

	// [3] --range_loop--
	// for index, value := range months {
	// 	fmt.Printf("bulan ke-%v adalah %v \n", index, value)
	// }
	// jika ada khasus kita tidak ingin menggunakan index maka bisa dilakukan cara dibawah
	// for _, value := range months {
	// 	fmt.Printf("bulan ini adalah %v \n", value)
	// 	value = "tahun"
	// }
	// // desc: index baru diinisasi, value sebagai variable yang direpresentasikan dalam for mengganti months
	// fmt.Println(months)


}