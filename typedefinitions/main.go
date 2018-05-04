package main

import "fmt"

type Test struct         { N int }
func (m *Test) Name()    { fmt.Println("abc")}


// NewTest does not inherit any functions of Test
// can access fields
type NewTest Test


// New2Test is composite type, it inherit all functions of Test
// can access fields
type New2Test struct {
	Test
}

// if embedded type is pointer you must initialized it
type New3Test struct {
	*Test
}


func main() {
	n := NewTest{}
	n.N = 1
	// n have no method
	// n.Name() // error
	v := (*Test)(&n)
	v.Name()

	v2 := Test(n)
	v2.Name()

	n2 := New2Test{}
	n2.N = 2
	n2.Name()

	n3 := New3Test{Test: new(Test)}
	// access filed N will panic if you do not initialized *Test
	n3.N = 3
	n3.Name()
}
