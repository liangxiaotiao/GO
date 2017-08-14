package main

import (
	"fmt"
)

func main() {
	//lxt_func1()

	////匿名函数
	//a := func() {
	//	fmt.Println("你是猪吗")
	//}
	//a()

	////闭包
	//func_1 := lxt_func2(10)
	//fmt.Println(func_1(1))
	//fmt.Println(func_1(2))
	//
	////panic and recover
	//lxt_Panic()
	////defer 操作
	//lxt_defer()

	lxt_func_homework()
}

func lxt_func1() {
	a := 1
	fmt.Println(a)
}

/**
	TODO 闭包模式 作用就是搞定匿名函数
 */
func lxt_func2(x int) func(int) int {
	//获取X的内存地址
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

//defer操作
func lxt_defer() {
	//最简单的操作,defer是由先进后出,由下至上的方式调用
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")

	for i := 0; i < 3; i++ {
		defer fmt.Println("DEFER", i)
	}
}

//异常处理机制
func lxt_Panic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover IN lxt_Recover")
		}
	}()
	//defer recover必须在Panic之前调用,不然在panic这一步错误以后就不会继续往下执行了
	panic("Panic in lxt_Painc")
}

/**
	这一块代码运行结果,因为defer是先进后出的机制,
	所以在for里面defer func() { fmt.Println("closuer i = ", i) }()这一句代码执行一次
	在轮到defer fmt.Println("defer i = ", i)执行一次

	最后调用fs报错是因为他在第一个for语句之后调用执行
	closuer i =  4
	defer i =  3
	closuer i =  4
	defer i =  2
	closuer i =  4
	defer i =  1
	closuer i =  4
	defer i =  0
	panic: runtime error: invalid memory address or nil pointer dereference
	[signal 0xc0000005 code=0x0 addr=0x0 pc=0x488c46]


 */
func lxt_func_homework() {
	var fs = [4]func(){}


	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("closuer i = ", i) }()
	}
	for _,f := range fs{
		f()
	}
}
