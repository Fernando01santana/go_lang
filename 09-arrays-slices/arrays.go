package main

import "fmt"

func main() {
	fmt.Println("Arrays e slice")

	var array1 [5]int
	fmt.Println(array1)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

	slice = append(slice, 10)
	fmt.Println(slice)

	//arrays internos

}
