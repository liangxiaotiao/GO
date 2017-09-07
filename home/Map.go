package main

import (
	"fmt"
	"sort"
)

func main() {
	//lxt_map1()

	//lxt_map2()

	//lxt_map3()

	//lxt_map4()
	homework()
}

func homework() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f"}
	s := make([]int, len(m))
	i := 0
	s1 := make([]string, len(m))
	//value := "a"
	for k, v := range m {
		s[i] = k
		s1[i] = v
		i++
	}
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(m)
	//把m当中map的key和value兑换 复制到m2当中
	m2 := map[string]int{}
	for k, v := range m {
		m2[v] = k
	}
	fmt.Println(m2)
}

func lxt_map4() {
	//定义个map
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f"}
	//定义一个slice
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)
}

func lxt_map3() {
	//这一代码就相当于初始化了一个slice的map
	Slice_map := make([]map[int]string, 5)
	for i := range Slice_map {
		//为slice的i进行初始化
		Slice_map[i] = make(map[int]string, 1)
		Slice_map[i][1] = "OK"
		fmt.Println(Slice_map[i])
	}
	fmt.Println(Slice_map)

}

func lxt_map2() {
	//多重map
	var Map_a map[int]map[int]string
	Map_a = make(map[int]map[int]string)
	//每一级的map都需要给他进行初始化 make操作,不然取不到地址
	Map_a[1] = make(map[int]string)
	Map_a[1][1] = "OK"
	a := Map_a[1][1]
	fmt.Println(a)

	b, ok := Map_a[2][1]
	if !ok {
		Map_a[2] = make(map[int]string)
	}
	//赋值必须在初始化完成后再赋值
	Map_a[2][1] = "GOOD"
	b = Map_a[2][1]
	fmt.Println(b, ok)
}

func lxt_map1() {
	var Map_a map[int]string
	Map_a = make(map[int]string)
	Map_a[1] = "OK"
	a := Map_a[1]
	fmt.Println(a)
}
