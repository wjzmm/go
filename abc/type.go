package main

import "fmt"

const (
	x = iota
	y
	z
)

func main() {
	var a = 2
	var b = 3
	c := 20
	var str string = "hello golang"
	d := false
	fmt.Println(a, b, c, d, str)
	fmt.Println(x, y, z)
}
