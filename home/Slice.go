package main

import (
	"fmt"
)

func main() {

	//stringbf()
	//Append()
	//Copy()
	Homework()
}

func Homework() {

	slice_a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	/*
		slice_b := make([]int,len(slice_a))
		copy(slice_b,slice_a[:len(slice_a)])
		这种方法相比较是多做了几部操作,首先我把slice_a的数组类型和长度赋值给了slice_b,然后以相同的长度下的slice作为copy

		slice_b := slice_a[:]
		这种方法是最简单的完整的把slice_a的所有赋值给了slice_b

		TODO 总结这两种赋值方式都可行,第一种过于讲究程序的步骤,而第二种方法比较直接一点
	 */
	slice_b := make([]int, len(slice_a))
	copy(slice_b, slice_a[:len(slice_a)])
	fmt.Println(slice_b)
}

func Copy() {
	slice_a := []int{1, 2, 3, 4, 5, 6}
	slice_b := []int{7, 8, 9}

	/*
	1.copy函数,把当前slice_b的数值copy到slice_a的数值当中,因为slice_b的长度只有三位,那么就相应的替换掉了slice_a的三个元素值
	2.相应的 如果是吧slice_a的数值copy到slice_b当中时,因为slice_b的长度只有三位,那么就只会copy三个元素值到slice_b当中
	3.同理,如果前者的长度大于后者的长度,那么会把后者所有的元素值都copy到前者当中去
	copy(slice_a,slice_b)
	还有一种形式就是,如果我只想copy指定的元素值时,那么我们可以用到
	copy(slice_a,slice_b[0:1])
	另外如果我想指定copy的元素值替换的元素值时,那么我们可以用到
	copy(slice_a[3:4], slice_b[0:1])
	 */
	copy(slice_a[3:4], slice_b[0:1])

	fmt.Println(slice_a)
}

func Append() {
	a := [] int{1, 2, 3, 4, 5, 6}
	slice_a := a[2:5]
	slice_b := a[0:3]
	fmt.Println(slice_a, slice_b)
	//append 如果slice的容量在append以后超出了当前slice的容量,那么会重新分配一个新的内存地址,不会影响到当前slice的内存地址
	slice_b = append(slice_b, 1, 1, 2, 4, 5, 5)
	//如果在改变了当前的slice的时候,会把当前所有slice依赖的内存都会改变
	slice_a[0] = 8
	fmt.Println(slice_a, slice_b)

}

func slice() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	s1 := a[:5] //这句代码就相当于从最开始取值到第五个元素,
	fmt.Println(s1)
	s2 := a[5:] //这句代码就相当于从第五个元素开始取值到最后,
	fmt.Println(s2)
	s3 := a[4:len(a)] //这句代码就相当于从第四个元素开始取值,到a的最大长度
	fmt.Println(s3)
}

func stringbf() {
	a := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	fmt.Println(a)
	slice_a := a[2:5]
	fmt.Println(slice_a)
	slice_b := a[5:len(a)]
	fmt.Println(slice_b)
	//len代表着这个slice长度,cap代表着这个slice的容量
	fmt.Println(len(slice_b), cap(slice_b))
	slice_c := slice_b[1:3]
	fmt.Println(slice_c)

}
