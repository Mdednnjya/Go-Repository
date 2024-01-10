package main

import (
	"fmt"
	"sort"
	// "strings"
)

func main(){
	// mendefinisikan variable
	// var name string = "Made Dananjaya"
	// var lorem string = "kata1%kata2%kata3"
	// var angka = []int{12, 12, 14, 100, 15, 1, 19}
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
        "zecember",
    }

			//strings 

	// --turns to lower case--
	// fmt.Println(strings.ToLower(name))

	// --checking string in variable--
	// fmt.Println(strings.Contains(name, "ibot"))
	// if false is returned, does mean the variabel dont contain the string ("ibot")

	// --replace some string--
	// fmt.Println(strings.ReplaceAll(name, "Made", "Komang"))

	// --finding index base on first letter
	// fmt.Println("String \"aya\" akan ditemukan pada index ke-", strings.Index(name, "aya"))

	// --slice elements--
	// fmt.Println(strings.Split(lorem, "1"))
	// splited := strings.Split(lorem, "%")
	// fmt.Println(splited[0])
	// fmt.Println(splited[1])
	// fmt.Println(splited[2])
	// menjadikan string menjadi array yang setiap index dipisahkan sesuai parameter ("%")


			//sort library
	
	// --sorting--
	// postifNum := []int{12, 9, 5, 4, 13, 17,1, 9}
	// fmt.Println("sebelum:")
	// fmt.Println(postifNum)

	// sort.Ints(postifNum)
	// fmt.Println("sesudah:")
	// fmt.Println(postifNum)

	// --Search value of int base on index--
	// index := sort.SearchInts(angka, 15)
	// fmt.Print("angka \"15\" berada pada index ke-",index)

	// --Search value of string base on index--
	// searched:= sort.SearchStrings(months, "december")
	// setelah diurutkan berdasarkan ASC maka urutan awal pada slice akan berubah
	// fmt.Print(searched)

}

