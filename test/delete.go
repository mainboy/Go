package main

import (
	"fmt"
)

func main() {
	str := make(map[int]string)
	count := 0
	str[count] = "hello world"
	count++
	str[count] = "string"
	fmt.Println(str)
	delete(str, "string")
	fmt.Println(str)
	fmt.Println(count)

}
