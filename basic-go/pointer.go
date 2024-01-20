package main

import (
	"fmt"
)

// sebuah variable pada golang akan disimpan kesebuah memory ex. "0xc000022070"

func updatedName(x *string){
	*x = "Walters"
	// ketika parameter masuk ke func maka didalamnya akan diganti seperti diatas
}

func main(){
	
	// --example-1--
	name := "Tifa Lockhart"

	// fmt.Printf("var value %v, stored in %v \n", name, &name)
	// // untuk melihat memori pada sebuah var maka perlu ditambahkan *&* sebelum nama var

	// --example-2--
	duplicatename := &name
	// var "duplicatename" akan menjadi pointer jadi var "name" dikarenakan kita memasukan memori var name kesana

	// fmt.Printf("Var value %v, stored in %v \n", *duplicatename, duplicatename)
	// secara otomatis "duplicatename" akan return nilai dari var "name", jika value name diganti maka nilai dari var "duplicatename" akan ikut terganti
	// menggunakan "*" sebelum nama var yang di duplicate untuk meng-print value dari variable, jika tidak maka akan direturn alamat memori

	// --example-3--
	name = "name changed"
	// fmt.Printf("nilai var terganti \n")
	// mengganti value
	// fmt.Printf("Var value %v, stored in %v \n", *duplicatename, duplicatename)
	// bisa dilihat jika nilai dari var yang dipointerkan diganti maka value dari var yang menjadi pointer juga akan terganti

	// -example-4--
	fmt.Printf("var \"name\" value is: %v\n", name)
	updatedName(duplicatename)
	fmt.Printf("after value change through func it's become\n")
	// mencoba mengganti value dari "name" tapi tidak spontan tetapi menggunakan func
	fmt.Printf("var \"name\" value is: %v\n", name)
	// value dari name bisa diganti, dikarenakan bantuan func yang dimana menerima parameter "*var"



}