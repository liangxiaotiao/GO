package main

import "fmt"

func main() {
	fmt.Println("你好啊")
	a := [...]int{9, 5, 2, 6, 8}
	fmt.Println(a)

	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] > a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp

			}
		}
	}
	fmt.Println(a)

	b := [...]int{9, 1, 5, 2, 8}

	fmt.Println(b)
	num1 := len(b)
	for i := 0; i < num1; i++ {
		for j := i; j < num1; j++ {
			if b[i] < b[j]+1 {
				temp := b[j]
				b[j] = b[i]
				b[i] = temp
			}
		}
	}
	fmt.Println(b)

}
