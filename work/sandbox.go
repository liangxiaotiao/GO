package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is ", time.Now())
	fmt.Println("My favorite number is ", rand.Intn(10))
	fmt.Println("Now you have %g problems. ", math.Nextafter(2, 6))
	fmt.Println("Addreturn number", add(100, 50))
	a, b := swap("龙思婷", "梁小跳")
	fmt.Println("I love all for you", a, b)
}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
