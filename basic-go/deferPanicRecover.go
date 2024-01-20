package main

import (
	"fmt"
	// "math"
)
// Trivia: golang tidak mengenal Try-Catch maka dikenalkannya lah dengan Defer, Panic, recover untuk eror handling.

// func logging(){
// 	fmt.Println("function Complete")
// }

// 	// [1] Defer
// func circleArea(r float64) float64{
// 	defer logging()
// 	// akan memberikan tanda jika suatu func selesai di ekseksi, ditaruh di baris pertama didalam func

// 	return math.Pi * r * r

// }

//  // [2] Panic - func yang digunakan untuk menghentikan program
// saat panic func berjalan dan program sudah terhenti, defer Func tetap akan dieksekusi
// func EndApp(){
// 	fmt.Println("Function Executed")
// }

// func RunApp(error bool){
// 	defer EndApp()

// 	if error == true {
// 		panic("Problem Found")
// 	}
// }

//  // [3] Recover - menghandle jika terdapat Panic func
// func EndApp(){
// 	fmt.Println("Function Executed")
// 	Message := recover()
// 	fmt.Printf("Status: %v \n", Message)
// 	// func recover() dipanggil di fungsi yang dijadikan defer dikarenakan pada dasarnya
// 	// defer tetap akan berjalan walaupun menemukan eror
// }

// func RunApp(error bool){
// 	defer EndApp()

// 	if error == true {
// 		panic("Problem Found")
// 	}
// 	// func recover tidak ditaruh setelah panic() karena otomatis tidak akan ter-run
// }


func main(){
	// 1. Defer
	// result := circleArea(12)
	// fmt.Println(result)

	// 2. Panic
	// RunApp(true)

	// 3. Recover
	RunApp(true)
	fmt.Println("Code Program terus berjalan karna Func Recover() bekerja")
}