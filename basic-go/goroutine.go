package main

import (
    "fmt"
    "time"
)

func f1() {
    fmt.Println("f1 started")
    time.Sleep(2 * time.Second)
    fmt.Println("f1 done")
}

func f2() {
    fmt.Println("f2 started")
    time.Sleep(3 * time.Second)
    fmt.Println("f2 done")
}

func f3() {
    fmt.Println("f3 started")
    time.Sleep(1 * time.Second)
    fmt.Println("f3 done")
}

func main() {
    f1()
    go f2()
    f3()
    time.Sleep(4 * time.Second)
}
