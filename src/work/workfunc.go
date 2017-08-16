package main

/*
函数被调用的基本格式如下：

pack1.Function(arg1, arg2, …, argn)
Function 是 pack1 包里面的一个函数，括号里的是被调用函数的实参（argument）：这些值被传递给被调用函数的形参（parameter，参考 第 6.2 节）。
函数被调用的时候，这些实参将被复制（简单而言）然后传递给被调用函数。
函数一般是在其他函数里面被调用的，这个其他函数被称为调用函数（calling function）。
函数能多次调用其他函数，这些被调用函数按顺序（简单而言）执行，理论上，函数调用其他函数的次数是无穷的（直到函数调用栈被耗尽）。
 */
import (
	"fmt"
	"math"
	"errors"
)

func main() {
	fmt.Println("GO函数 学习")
	//a, b, c := lxtFuncReturn1(20, 10)
	//fmt.Println("和:", a, " 积:", b, " 差:", c)
	//MySqrt()
	n:=0
	reply := &n
	Multiply(10,5,reply)
	fmt.Println("Multiply",*reply)
}

//编写一个函数，接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值。
func lxtFuncReturn(i, j int) (int, int, int) {
	return i + j, i * j, i - j
}
func lxtFuncReturn1(i, j int) (a, b, c int) {
	a = i + j
	b = i * j
	c = i - j
	return a, b, c
}
func MySqrt() {
	num1, err := MySqrt1(-0.5)
	fmt.Println(num1, err)
	num2, err := Mysqrt2(1)
	fmt.Println(num2, err)
}

//编写一个名字为 MySqrt 的函数，计算一个 float64 类型浮点数的平方根，如果参数是一个负数的话将返回一个错误。编写两个版本，一个是非命名返回值，一个是命名返回值。
func MySqrt1(f float64) (float64, error) {
	if f < 0 {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}
	return math.Sqrt(f), nil
}

func Mysqrt2(f float64) (ret float64, err error) {
	if f < 0 {
		//then you can use those variables in code
		ret = float64(math.NaN())
		err = errors.New("I won't be able to do a sqrt of negative number!")
	} else {
		ret = math.Sqrt(f)
		//err is not assigned, so it gets default value nil
	}
	return
}
// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}



