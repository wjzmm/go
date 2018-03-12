package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false

	if a && b {
		fmt.Printf("第一行 - 条件为true\n")
	}

	if a || b {
		fmt.Printf("第二行 - 条件为true\n")
	}

	a, b = b, a
	if a && b {
		fmt.Printf("第三行 - 条件为true\n")
	} else {
		fmt.Printf("第三行 - 条件为false\n")
	}

	if !(a && b) {
		fmt.Printf("第四行 - 条件为true\n")
	}

}
