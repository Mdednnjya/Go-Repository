package main

import (
	"fmt"
)

// ket: nil (data kosong) hanya dapat diinisiasi dengan jika menggunakan dengan interface, function, map, slice, pointer, & channel

func myName (name string) map[string]string{
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main(){
	// declaration
	var scanner string

	// melakukan input variable dari scanner
	fmt.Print("data Nama: ")
	_, err:= fmt.Scanln(&scanner)
	if err != nil {
		fmt.Println(err)
		return
	}
	// memasukan nilai variable ke func
	data := myName(scanner)

	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Printf("Data adalah %v \n", data)
	}
	// return an nil value



}