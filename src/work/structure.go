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
	//lxtSwitchWork3()
	//lxtForRange()
	//lxtForWork()
	//lxtBreakAndContinue()
	lxtGotolevel()
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

func lxtForRange() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
	fmt.Println()
	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}

func lxtForWork() {
	//应该是循环打印5个0 v在循环中赋值为5后重新进入循环体，又被重新初始化了
	//for i := 0; i < 5; i++ {
	//	var v int
	//	fmt.Printf("%d ", v)
	//	v = 5
	//}
	//应该是个无限循环，因为这段代码没有条件，所以无限循环
	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//}
	//（打印一个0 因为他没有修饰语句，所以成立一次后结束）
	// 結論错误!!!经过运行，得出因为没有修饰语句，此段代码每次都成立i<3，所以为无限循环
	//for i := 0; i < 3;{
	//	fmt.Printf("\nValue of i: %d", i)
	//}
	//应该是循环五次,因为在循环尾部s = s + "a"这句代码相当于为上面的变量赋值,累加五次后与条件成立退出循环
	s := ""
	for ; s != "aaaaa"; {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}

	//循环3次,等到i不满足<3的条件时 退出循环
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}

//break\continue的使用
func lxtBreakAndContinue() {
	//关键字 continue 忽略剩余的循环体而直接进入下一次循环的过程，但不是无条件执行下一次循环，执行之前依旧需要满足循环的判断条件。
	for i := 1; i < 10; i++ {
		if i == 5 {
			continue
		}
		print(" ", i)
	}
	//因此每次迭代都会对条件进行检查（i < 0），以此判断是否需要停止 循环。如果退出条件满足，则使用 break 语句退出循环。
	//一个 break 的作用范围为该语句出现后的最内部的结构，它可以被用于任何形式的 for 循环（计数器、条件判断等）。
	// 但在 switch 或 select 语句中（详见第 13 章），break 语句的作用结果是跳过整个代码块，执行后续的代码。
	//	下面的示例中包含了嵌套的循环体（for4.go），break 只会退出最内层的循环：
	for i := 1; i < 10; i++ {
		if i == 5 {
			break
		}
		println()
		print(" ", i)
	}
}

//label和goto的使用
func lxtGotolevel() {
	//for、switch 或 select 语句都可以配合标签（label）形式的标识符使用，即某一行第一个以冒号（:）结尾的单词（gofmt 会将后续代码自动移至下一行）。

LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}
