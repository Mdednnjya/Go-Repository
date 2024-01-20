package main

import (
	"fmt"
)

// struct is just like a custome data type

// type Animal struct{
// 	name, habitat string
// 	maxAge int
// 	// inisiasi variable struct biasanya diawali dengan huruf besar aka "Pascal Case"
// }

// func (animal Animal) sayhelo(){
// 	fmt.Printf("got hello from %v \n", animal.name)
// }
// // funct disepesifikan untuk struct apa pada parameter pertama "([nama struct] [type struct])"

// func (animal Animal) addAge(n int) int {
// 	return animal.maxAge + n
// }

func main(){
	// --declaration struct--
	// // 1. dengan inisiasi Var Animal
	// var lion Animal
	// lion.maxAge = 12
	// lion.name = "Ruco"
	// lion.habitat = "forest"

	// // 2. inisiasi var sekaligus atribut
	// tiger := Animal{
	// 	name: "Gia",
	// 	habitat: "forest Asia",
	// 	maxAge: 22,
	// }

	// // 3. lebih cepat, sesuai urutan
	// crocodile := Animal{"bla", "Swep", 23}
	

	// // output
	// fmt.Println(lion.name)
	// fmt.Println(tiger)
	// fmt.Printf("nama buaya adalah %v \n", crocodile.name)

	// --method in struct--
	// example-1
	// lion := Animal{"Lion-frezzy", "Canyon", 12}
	// lion.sayhelo()
	// fmt.Printf("%v age is %v", lion.name, lion.addAge(12))



}