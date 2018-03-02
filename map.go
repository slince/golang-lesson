package main

import "fmt"


func main(){


	m := make(map[string]int)

	m["string"] = 123

	fmt.Println(m)


	type Vertical struct{
		x string
		y int
	}


	var position map[string]Vertical

	position = make(map[string]Vertical)

	position["a"] = Vertical{
		"hello",
		50,
	}

	fmt.Println(position)



	var zuobiao = map[string]string{
		"hello": "world",
	}







	fmt.Println(zuobiao)


	fmt.Println(zuobiao["hello"])
	fmt.Println(zuobiao["hell"])




   var value, exists = zuobiao["hello"]

	fmt.Println(value, exists)

	delete(zuobiao, "hello")

	var value2, exists2 = zuobiao["hello"]
	fmt.Println(value2, exists2)
}
