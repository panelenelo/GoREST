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

func test3() {
	intSet := make(map[int]bool)
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	// fmt.Println(len(vals), len(intSet))
	// fmt.Println((intSet[5]))
	// fmt.Println(intSet[500])
	// if intSet[100] {
	// 	fmt.Println("100 is in the Set")
	// }
	if _, ok := intSet[2]; ok {
		fmt.Println(ok)
	}
}

func test4() {
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	sub1 := greetings[:2]
	sub2 := greetings[1:4]
	sub3 := greetings[3:5]
	fmt.Println(sub1)
	fmt.Println(sub2)
	fmt.Println(sub3)
}

func test5() {
	var message = "Hi 👩 and 👨"
	runes := []rune(message)
	fmt.Printf("%c \n", runes[3])
}

func test6() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	pedro := Employee{
		"pedro",
		"mango",
		12,
	}

	joao := Employee{
		"joao",
		"manuel",
		31,
	}

	var carlos Employee
	carlos.firstName = "carlos"
	carlos.lastName = "emanuel"
	carlos.id = 51

	fmt.Println(carlos, pedro, joao)
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func main() {
	test6()
}
