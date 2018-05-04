package main

import "fmt"



func main() {

	n1 := new(int)
	fmt.Println(n1) // console the address of n1
	fmt.Println(*n1 == 0) // zero value

	type T struct {
		I int
		A string
		Next *T
	}
	n2 := new(T)
	fmt.Println(n2)
	fmt.Println(n2.I == 0)
	fmt.Println(n2.Next == nil)
	fmt.Println(n2.A == "")

	n3 := new([]int)
	fmt.Println(n3)
	fmt.Println(*n3 == nil)



	//---------
	m1 := make([]int, 1)
	fmt.Println(m1)
	m2 := make(map[int]string)
	m2[0] = "abcde"

	m3 := make(chan bool)
	m4 := make(chan bool, 5)

	_ = len(m3)
	_ = len(m4)
}


