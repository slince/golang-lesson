package main

import "fmt"

func main() {
	type Vertex struct {
		X int
		Y int
	}
	fmt.Println(Vertex{1, 2})


	hello := Vertex{1, 2}

	fmt.Println(hello.X);

	//p := &hello
	hello.X = 1e9
	fmt.Println(hello)
}
