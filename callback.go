package main


import (
	"fmt"
	"math"
)


func returnCallback() func(int) int{

	return func(x int) int {
		return x
	}
}

func main(){

	func1 := func(x, y float64) float64{
		return math.Pow(x, y);
	}
	fmt.Println(func1(2,3))



}