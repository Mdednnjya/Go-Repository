package main

import (
	"fmt"
)

func emptyInterface() interface{}{
	return "ok"
}

func main(){
	value := "12"
	// resultValue := value.(int)
	
	fmt.Println(value)
}