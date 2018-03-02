package main

import "fmt"


func add(x, y int) int {
	return x+y
}

func multiReturn(x, y int)(int, int, int){
	var z = x + y
	return x, y ,z
}

func namedReturn(x, y int)(z, g int){

	z = x + y
	g = x -y
	return
}

func main(){
	fmt.Println(add(1, 2))

	fmt.Println(multiReturn(1, 2))

	fmt.Println(namedReturn(1, 2))
}


