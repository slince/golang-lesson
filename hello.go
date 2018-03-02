package main

import "fmt"

import (
	"runtime"
)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}


var x, y int = 1 , 2

var (
	a string = "hello"
	b string = "world"
)
func main() {
	k := 3
	//fmt.Println(split(17))
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(k)
	fmt.Println(a)
	fmt.Println(b)

	a := string(b)
	fmt.Println(a)

	i :=5;
	for ;i<=10; {
		i++
		fmt.Println(i)
	}

	if true {
		fmt.Println(true)
	}

	fmt.Println(runtime.GOOS)

	switch runtime.GOOS{
	    case "windows":
		    fmt.Println("window")
			break
	}

	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")


	var hello *string;

	hello = &a

	fmt.Println(*hello)
	*hello = "world"
	fmt.Println(a)
}
