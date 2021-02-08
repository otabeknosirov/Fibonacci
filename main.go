package main

import "fmt"

func main() {

	var in int 
	fmt.Println("Enter the fibannacci number...")
	fmt.Scan(&in)

	res := fib(in)

	fmt.Println(res)
}

func fib(fibNum int) int{
 
	if(fibNum == 0) {
		return 0
	}

	if(fibNum == 1) {
		return 1
	}
	return fib(fibNum - 1) + fib(fibNum - 2)
}