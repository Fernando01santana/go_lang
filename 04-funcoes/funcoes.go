package main

import "fmt"

func calculosMatematicos(x int8, y int8) (int8, int8) {
	soma := x + y
	subtracao := x - y

	return soma, subtracao
}

func main() {
	soma, _ := calculosMatematicos(2, 4)
	fmt.Println(soma)
}
