package main

import (
	"fmt"
)

type D struct {
	Name string
}
type B struct {
	Name string
}
type C int

func main() {

	//lxtMethod1()
	lxtMethodWork()
}

func lxtMethod1() {
	a := D{}
	a.print()
	fmt.Println("获取", a.Name)

	b := B{}
	b.print()
	fmt.Println("获取", b.Name)
}

func lxtMethodWork() {
	var c C
	c.print(100)
	fmt.Println(c)
}

func (a *D) print() {
	a.Name = "你好"
	fmt.Println("打印", a)
}
func (b B) print() {
	b.Name = "龙骚骚"
	fmt.Println("打印", b)
}

func (c *C) print(num int) {
	*c += C(num)
}
