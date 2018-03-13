package main

import "fmt"

func main() {
	var my_map map[string]string
	my_map = make(map[string]string)
	my_map["x"] = "xxxx"
	my_map["y"] = "yyyy"
	my_map["z"] = "zzzz"

	for key := range my_map {
		fmt.Println("value of", key, "is", my_map[key])
	}

	skey, ok := my_map["x"]

	if ok {
		fmt.Println("ok, key is", skey)
	} else {
		fmt.Println("not ok")
	}

	delete(my_map, "z")
	fmt.Printf("删除元素后\n")

	for key := range my_map {
		fmt.Println("value of", key, "is", my_map[key])
	}

}
