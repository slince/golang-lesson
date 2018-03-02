package main


import "fmt"

func main(){

	//固定长度，默认0 填充
	var x [10]int

	for i := 0; i<10; i++ {
		x[i] = i + 1
	}
	fmt.Println(x)

	//数组赋值
	x = [10]int{2,2,3,4,5,6,7,7,8,9}

	fmt.Println(x)

	//声明时赋值
	y := []int{1,2,3,4}
	fmt.Println(y)


	fmt.Println(len(y))
	fmt.Println(len("hello world"))

	//数组裁剪
	fmt.Println(y[0:2])

	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
