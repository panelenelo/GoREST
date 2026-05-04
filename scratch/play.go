package main

import (
	"fmt"
)

func test1() {
	var i int = 20
	var f float32 = float32(i)
	fmt.Printf("i = %d\n", i)
	fmt.Printf("f = %f\n", f)
	//fmt.Println("i = ", i, "\nf = ", f)
}

func test2() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	println(b)
	println(smallI)
	println(bigI)
	println("--------")
	println(b + 1)
	println(smallI + 1)
	println(bigI + 1)

}

func main() {
	test2()
}
