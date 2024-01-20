package main

import (
	"fmt"
)


func changeName(n string) {
	n = "wellb"
	// n pada parameter fungsi changeName adalah clone/salinan yang jika nilainya diubah tidak akan berpengaruh ke variable asli
}

func main(){
	// 1. Concept pass By Value adalah dimana ketika kita memberikan nilai dari variabel a ke variabel b,
	// maka ketika kita memodifikasi nilai a maka tidak akan berpengaruh ke variable b
	// Hal ini terjadi karena nilai variable b dialokasikan ke lokasi memori lain (tidak sama dengan a)

	// declare
	name := "tifa"

	changeName(name)
	// mencoba mengganti value dari variable name dengan bantuan func, tetapi tidak akan berhasil dikarenakan yang di modifikasi adalah clonenya

	fmt.Println(name)
	// output
}