package main


import "fmt"


type A struct{
	a int
	b string
	c *int
}

func main(){

	//var a, b, c int = 1, 2,2
	b, c, d := 1, 2 ,2
	e := "hello"

	var a = A{1,e, &c}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(*a.c)


	var x [10]int
	fmt.Println(x)

}