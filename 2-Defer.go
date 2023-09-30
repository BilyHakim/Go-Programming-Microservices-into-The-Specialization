package main

import "fmt"

//func main() {
//	defer fmt.Println("defer function starts to execute")
//	fmt.Println("Hai Everyone")
//	fmt.Println("Welcome back to Go Learning Center")
//}

func main() {
	callDeferFunc()
	fmt.Println("Hai everyone!")
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func start to execute")
}
