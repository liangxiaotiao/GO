package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println("控制结构学习")
	//lxtIfElse(11, 12)
	//lxtIfElse2(11, 22)
	//lxtIfElse3()
	//lxtIfElse4()
	//lxtSwitch1(5)
	//season(8)
	//lxtFoWork1()
	//lxtFoWork2()
	lxtSwitchWork3()
}

//记录 %d是format整数中添加
func lxtIfElse(a, b int) {
	if a > b {
		fmt.Printf("a = %d >b = %d ?yes\n", a, b)
	} else {
		fmt.Printf("a = %d >b = %d ?no\n", a, b)
	}
}

func lxtIfElse2(first, cond int) {
	if first <= 0 {
		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {
		fmt.Printf("first is between 0 and 5\n")
	} else {
		fmt.Printf("first is 5 or greater\n")
	}
	if cond = 5; cond > 10 {
		fmt.Printf("cond is greater than 10\n")
	} else {
		fmt.Printf("cond is not greater than 10\n")
	}
}

func lxtRerutn(value int) int {
	return value - 5
}
func lxtRerutnString(value string) (name string) {
	return value
}
func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	} // error case
	return math.Sqrt(f), true
}

func lxtIfElse3() {
	var max int = 20
	if value := lxtRerutn(10); value > max {
		fmt.Printf("lxtIfElse3 value = %d >max = %d :true", value, max)
	} else {
		fmt.Printf("lxtIfElse3 value = %d >max = %d :false", value, max)

	}
}

func lxtIfElse4() {
	var orig string = "123"
	// var an int
	var newS string
	// var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// anInt, err = strconv.Atoi(origStr)
	//把字符串orig转出int类型,然后去判断是否转换成功,err返回true\false
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		//遇到错误跳出 不执行下面的语句了
		return
	}
	value, err1 := strconv.Atoi(lxtRerutnString("321"))
	if err1 != nil {
		fmt.Printf("An error occured in pack1.Function1 with parameter %v\n", value)
		//遇到错误时结束程序
		os.Exit(1)
		return
	}
	// 未发生错误，继续执行：
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

	//如果您像下面一样，没有为多返回值的函数准备足够的变量来存放结果：
	t, ok := mySqrt(25.0)
	if ok {
		fmt.Println("T-", t)
	}
}

func lxtSwitch1(number int) {
	switch number {
	case 1:
		fmt.Println("This number is equls to", number)
	case 5:
		fmt.Println("This number is equls to", number)
	default:
		fmt.Println("Default This number is equls to", number)
	}

	switch value := lxtRerutn(20); {
	case value == 1:
		fmt.Println("This number is equls to", value)
	case value == 10:
		fmt.Println("This number is equls to", value)
	default:
		fmt.Println("Default This number is equls to", value)
	}

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

}

//作业 写一个 Season 函数，要求接受一个代表月份的数字，然后返回所代表月份所在季节的名称
func season(months int) {
	switch months {
	case 1, 2, 3:
		fmt.Println("春季")
	case 4, 5, 6:
		fmt.Println("夏季")
	case 7, 8, 9:
		fmt.Println("秋季")
	case 10, 11, 12:
		fmt.Println("冬季")
	default:
		fmt.Println("出错啦")
	}
}

func lxtFoWork1() {
	for i := 1; i < 16; i++ {
		fmt.Println("依次打印:", i)
	}

	i := 0
START:
	fmt.Println("The counter is", i)
	i++
	if i < 15 {
		goto START
	}
}
func lxtFoWork2() {
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			print("G")
		}
		println()
	}
	for i := 25; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			print("G")
		}
		println()
	}

}

const (
	FIZZ     = 3
	BUZZ     = 5
	FIZZBUZZ = 15
)

func lxtSwitchWork3() {
	//写一个从 1 打印到 100 的程序，但是每当遇到 3 的倍数时，不打印相应的数字，但打印一次 "Fizz"。遇到 5 的倍数时，打印 Buzz 而不是相应的数字。
	// 对于同时为 3 和 5 的倍数的数，打印 FizzBuzz
	for i := 1; i <= 100; i++ {
		switch {
		case i%FIZZBUZZ == 0:
			fmt.Println("打印 :FIZZBUZZ")
		case i%FIZZ == 0:
			fmt.Println("打印 :FIZZ")
		case i%BUZZ == 0:
			fmt.Println("打印 :BUZZ")
		default:
			fmt.Println("打印 :", i)
		}
	}

	//使用 * 符号打印宽为 20，高为 10 的矩形。
	w, h := 20, 10
	for y := 1; y <= h; y++ {
		for x := 1; x <= w; x++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
