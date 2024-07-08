package main

import "fmt"

func main() {
	fmt.Println("hello world")
	f := test(10)
	fmt.Println(f)

}

func test(x int) int {
	fmt.Println(x)
	return x + 10
}
