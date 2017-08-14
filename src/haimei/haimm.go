package main

import (
	"fmt"
	"unsafe"
)

const c = "C"
const LENGTH int = 10
const WIDTH = 5

var area int

const a1, a2, a3 = 1, false, "lst"
const (
	b1 = "abc"
	b2 = len(b1)
	b3 = unsafe.Sizeof(b1)
)
const (
	c1 = iota
	c2
	c3
	c4 = "哈哈哈"
	c5
	c6 = 100
	c7
	c8 = iota
	c9
)

var v int = 5
var q = "梁小跳"
var w string = "买车啦啦啦啦"
var e bool

var v1, v2, v3 int
var (
	v4 int
	v5 bool
)
var v6, v7 int = 1, 2
var v8, v9 = "哎哟", "不错哦"

type T struct {
}

func (t T) Method1() {

}
func init() {
	fmt.Println("张鹏蠢子")
}
func xixi() {
	fmt.Println("梁尚胖子")
}
func main() {

	var a int
	fmt.Println("资涛哑子")
	xixi()
	fmt.Println(a)
	fmt.Println(q, w, e)
	v10, v11 := "哎哟", "不错哦"
	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11)

	area = LENGTH * WIDTH
	fmt.Printf("面积:%d", area)
	fmt.Println()
	fmt.Println(a1, a2, a3)

	fmt.Println(b1, b2, b3)

	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9)

}
