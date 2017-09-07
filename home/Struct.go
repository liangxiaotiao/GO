package main

import (
	"fmt"
)

func main() {
	//lxt_Struct1()
	lxt_Struct3()
}

func lxt_Struct2() {
	p := &person{
		Name: "梁小跳",
		Age:  22,
	}
	fmt.Println(p)
	A(p)
	fmt.Println(p)
}
func A(person *person) {
	person.Age = 23
	fmt.Println("嘻嘻", person)
}

type person struct {
	Name string
	Age  int
}

func lxt_Struct1() {
	one := struct {
		Name string
		Age  int
	}{
		Name: "龙骚骚",
		Age:  22,
	}
	fmt.Println(one)
}

type people struct {
	Name string
	Age  int
	women struct {
		City, phone string
	}
}

func lxt_Struct3() {
	people := people{
		Name: "龙骚骚",
		Age:  22,
	}
	people.women.City = "耒阳"
	people.women.phone = "18607475151"
	fmt.Println(people)
}
