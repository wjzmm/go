package main

import "fmt"

func main() {
	var a int = 20
	var ip *int

	ip = &a

	fmt.Printf("a 的变量地址是: %x\n", &a)

	fmt.Printf("ip 变量的存储地址是: %x\n", &ip)

	fmt.Printf("*ip 变量的值：%d\n", *ip)
}
