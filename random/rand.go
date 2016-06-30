package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	times := int64(time.Now().Nanosecond())
	rand.Seed(times)

	fmt.Println("Rand int start!")
	for i := 0; i < 10; i++ {
		v := rand.Intn(50)
		fmt.Printf("%d\n", v)

	}

	fmt.Println("Rand float start!")
	for i := 0; i < 10; i++ {
		v := rand.Float32()
		fmt.Printf("%f\n", v)
	}

}
