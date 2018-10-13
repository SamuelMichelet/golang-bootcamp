package main

import (
	"fmt"
)

var y, x int
var s = "hy"
var bt = `Hy guys
	how are you?`

type ownType int

var b ownType

func main() {
	y = 2
	n, err := fmt.Println("hy guy")
	fmt.Println(n)
	fmt.Println(err)
	x = 42
	fmt.Println(x)
	fmt.Printf("type: %T", x)
	x = 12
	fmt.Println(x + y)
	fmt.Println(s)
	b = 77
	fmt.Println(b)
	fmt.Printf("type: %T, %v\n", b, b+2)
	bInt := int(b)
	fmt.Println(bInt)
	bs := string(b) // utf8
	fmt.Println(bs)
}
