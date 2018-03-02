package main

import "fmt"



func main(){
	//var y [8]int= [8]int{1, 2, 4, 8, 16, 32, 64, 128}
	//
	//var x []int

	//z := []int{12,3,4,5}
	//
	//x = append(x, 1)
	//x = append(x, 2, 3)



	z := make([]int, 5)
    z = append(z, 1, 2);

    for _, value := range z {
		//fmt.Println(key)
		fmt.Println(value)
	}

	fmt.Println(len(z))
	fmt.Println(cap(z))

	fmt.Println(uint64(10))


	//
	//fmt.Println(x)
	//fmt.Println(y)




}
