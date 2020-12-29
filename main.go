package main

import "fmt"

func Add(nums ...int) int {
	res := 0
	for n := range nums {
		res += n
	}
	return res
}

func main() {
	fmt.Println(Add(1, 2, 3))
}
