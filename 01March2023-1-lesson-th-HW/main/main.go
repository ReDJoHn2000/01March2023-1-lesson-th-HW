package main

import (
	"fmt"
)

func main(){

var firstNum float32
var secondNum float32

fmt.Print("Enter your first number: ")
fmt.Scanln(&firstNum)

fmt.Print("Enter your second number: ")
fmt.Scanln(&secondNum)

fmt.Println("The sum is = ", plus(firstNum, secondNum))
fmt.Println("The difference is = ", minus(firstNum, secondNum))
fmt.Println("The multiplication is = ", multp(firstNum, secondNum))
fmt.Println("The division is = ", divis(firstNum, secondNum))
}

func plus(x, y float32) float32{
	return x + y
}
func minus(x, y float32) float32{
	return x - y
}
func multp(x, y float32) float32{
	return x * y
}
func divis(x, y float32) float32{
	return x / y
}

