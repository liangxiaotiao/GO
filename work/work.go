package main

import (
	"fmt"
	"runtime"
	"os"
	"unicode/utf8"
	"strings"
	"strconv"
	"time"
)

//const 在函数外为全局常量 在函数体内则为局部常量
//var 定义变量 采用驼峰命名法

func main() {
	//a, b, c := lxt_work1(1, 2, 3)
	//fmt.Println(a, b, c)

	//lxtConstants()

	//lxtRuntime()
	//lxtCasting()
	//lxtLianxi()
	//lxtString()
	//lxtTime()
	lxtZhizhen()
}

func lxt_work1(one, two, three int) (one_1, two_1, three_1 int) {
	one_1 = one + 1
	two_1 = two + 1
	three_1 = three + 1
	return one_1, two_1, three_1
}

func lxtConstants() {

	const (
		//常量
		Pi = 3.1415926
		Wi = "你是谁呢"
		Xi = "嘻嘻嘻嘻"
	)
	const (
		//自动递增 iota
		A1 = iota
		B1
		C1
		D1 = "你是"
		E1 = iota
		F1
	)
	fmt.Println(Pi, Wi, Xi)

	fmt.Println(A1, B1, C1, D1, E1, F1)
}

func lxtRuntime() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}

func lxtCasting() {
	var n int16 = 4
	var m int32
	var s string
	m = int32(n)
	s = string(n)
	fmt.Println(n, m, s)
}
//操作字符串
func lxtLianxi() {
	// count number of characters:
	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
	fmt.Printf("The number of characters in string str2 is %d", utf8.RuneCountInString(str2))
}

func lxtString() {
	var str string = "The is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?\n ", str, "Th")
	//TODO 前缀和后缀
	//HasPrefix 判断字符串 s 是否以 prefix 开头：
	fmt.Printf("%t\n", strings.HasSuffix(str, "ng"))
	//HasSuffix 判断字符串 s 是否以 suffix 结尾：
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	//TODO 字符串包含关系
	//Contains 判断字符串 s 是否包含 substr：
	fmt.Printf("%t\n", strings.Contains(str, "lxt"))
	//TODO 判断子字符串或字符在父字符串中出现的位置（索引）
	//Index 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
	fmt.Printf("%t\n", strings.Index(str, "s"))
	//LastIndex 返回字符串 str 在字符串 s 中最后出现位置的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
	fmt.Printf("%t\n", strings.LastIndex(str, "s"))
	//Replace 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串，如果 n = -1 则替换所有字符串 old 为字符串 new：
	//TODO  字符串替换，统计字符串出现次数，重复字符串
	fmt.Printf("%t\n", strings.Replace(str, "The", "Lxt", 3))
	//Count 用於计算字符串str在字符串s当中出现的非重叠次数
	fmt.Printf("%t\n", strings.Count(str, "s"))
	//Repeat 用于重复 count 次字符串 s 并返回一个新的字符串：
	fmt.Printf("%t\n", strings.Repeat(str+"，  ", 3))
	//TODO 修改字符串大小写
	//ToLower 将字符串中的 Unicode 字符全部转换为相应的小写字符：
	fmt.Printf("%t\n", strings.ToLower(str))
	//ToUpper 将字符串中的 Unicode 字符全部转换为相应的大写字符：
	fmt.Printf("%t\n", strings.ToUpper(str))
	//TODO 修剪字符串
	//你可以使用 strings.TrimSpace(s) 来剔除字符串开头和结尾的空白符号；
	fmt.Printf("%t\n", strings.TrimSpace(str))
	// 如果你想要剔除指定字符，则可以使用 strings.Trim(s, "cut") 来将开头和结尾的 cut 去除掉,该函数的第二个参数可以包含任何字符，
	fmt.Printf("%t\n", strings.Trim(str, "The,string"))
	// 如果你只想剔除开头或者结尾的字符串，则可以使用 TrimLeft 或者 TrimRight 来实现。
	fmt.Printf("%t\n", strings.TrimLeft(str, "The"))
	fmt.Printf("%t\n", strings.TrimRight(str, "string"))
	//TODO 分割字符串
	//	strings.Fields(s) 将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块，并返回一个 slice，如果字符串只包含空白符号，则返回一个长度为 0 的 slice。
	fmt.Printf("%t\n", strings.Fields(str))
	//strings.Split(s, sep) 用于自定义分割符号来对指定字符串进行分割，同样返回 slice。
	fmt.Printf("%t\n", strings.Split(str, "s"))
	//TODO 字符串转换
	var num string = "123456"
	an, _ := strconv.Atoi(num)
	an += 5
	fmt.Printf("%t\n", an)
}

//时间
func lxtTime() {
	var week time.Duration
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())

	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(week)
	fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011

	// formatting times:
	fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
	fmt.Println(t.Format(time.ANSIC)) // Wed Dec 21 08:56:34 2011
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	// Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221
}

func lxtZhizhen()  {
	//&去内存地址的符号 *指向指针的符号
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)

	//这段代码就是把intP指向指针,但是取值的时候intP = &i1 又取到了i1的这个内存地址,所以第一个打印了内存地址,第二个intP重新指向指针打印了5
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p) // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s) // prints same string
}
