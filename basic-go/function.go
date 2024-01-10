package main

import (
	"fmt"
	"math"
	
)

// --Basic Function--
// example-1
// func turnEngine(car string){
// 	fmt.Printf("Engine %v: on", car)
// }

// example-2
// func multiplication (a float32, b float32){
// 	result :=  a * b
// 	fmt.Printf("the result of %v * %v is %v", a, b, result)
// }

// example-3
// func expression(car string){
// 	fmt.Printf("%v issa Hypercar \n", car)
// }

// example-4
// func repeated(element []string, funct func(string)){
// 	for _, value := range element {
// 		funct(value)
// 	}
// // menggunakan func "repeated" untuk merepetisi print dari func "expression" dengan nilai menyesuaikan dengan parameter 1(array)
// }

// --function with return value--

// example-1
// func circleArea(r float64) float64{
// 	return math.Pi * r * r
// }
// ketika akan mengreturn value maka disamping parameter harus ditulisakan tipe data yang ingin di return (Ex. float64)

// example-2


func main(){

	// declaration

	// --func--
	// turnEngine("Mclaren")
	// multiplication(2, 3)
	// repeated([]string{"Ferrari", "LaFerrari", "V12 Hybrid"}, expression)
	
	// example func with scanner
	// var input1 float64 
	// fmt.Printf("Input value for circle1: ")
	// fmt.Scanln(&input1)

	// var input2 float64 
	// fmt.Printf("Input value for circle2: ")
	// fmt.Scanln(&input2)

	// circle1Result := circleArea(input1)
	// circle2Result := circleArea(input2)
	// fmt.Printf("here is the result of %0.2f (circle 1) and %0.2f (circle 2)", circle1Result, circle2Result)


}
