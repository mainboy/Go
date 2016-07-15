package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	const str = "hello world\n askhdfka\n"
	const str1 = "hello world"
	s := bufio.NewReader(strings.NewReader(str))
	for {
		line, e := s.ReadBytes('\n')
		if e != nil {
			break
		}
		fmt.Printf("%s", line)
	}
	fmt.Printf("%s", str1)

}
