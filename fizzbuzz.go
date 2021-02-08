package main

import "fmt"

func main()  {
	var num int
	fmt.Println("Enter number for FizzBuzz...")
	fmt.Scan(&num)

	fizzBuzz(num)
}

func fizzBuzz(num int){

	for i:=1; i <= num; i++ { 
       if(i % 3 == 0 && i % 5 != 0){
		   fmt.Println("Fizz")
	   }
	   if(i % 5 == 0 && i % 3 != 0){
		   fmt.Println("Buzz")
	   }
       if(i % 3 == 0 && i % 5 == 0){
		   fmt.Println("FizzBuzz")
	   }
	}
}