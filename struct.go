package main


import "fmt"

func main()  {

	type 坐标 struct{
		a int
		b int
		c,d string
		e int
	}

	var a = 坐标{1,2, "hello", "world", 3,}
	var b 坐标 = 坐标{1,2, "hello", "world", 3,}

	fmt.Print(a)
	fmt.Print(b)
}