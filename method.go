package main


import "fmt"

type Vertex struct{
	x, y int
}


func (v *Vertex) Add() int {
	return v.x +v.y
}


func(v *Vertex) reduce(num int) int{
	return (v.x + v.y) * num
}

func main(){
	m := &Vertex{1,2}

	fmt.Print(m.Add())
	fmt.Print(m.reduce(2))
}



