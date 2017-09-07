package main

import (
	"fmt"
)

//定义一个接口
type USB interface {
	//接口里面包含两个方法
	Name() string
	Connecter
}

//嵌入接口
type Connecter interface {
	Connect()
}
type PhoneConnect struct {
	name string
}

func main() {
	lxtInterFace()
}

func (pc PhoneConnect) Name() string {
	return pc.name
}

func (pc PhoneConnect) Connect() {
	fmt.Println("PhoneConnect :", pc.name)
}

func lxtInterFace() {
	var usb USB
	usb = PhoneConnect{"lxtInterFace"} //为这个PhoneConnect里面的name赋值
	usb.Connect()
	DisConnect(usb)
	lxtSwitch(usb)
}

func DisConnect(usb USB) {
	//类型断言
	if pc, ok := usb.(PhoneConnect); ok {
		pc.name = "你好啊"
		fmt.Println("DisConnect :", pc.name)
		return
	}
	println("DisConnect")
}

//传进来一个空接口和switch/case的判断运用
func lxtSwitch(usb interface{}) {
	//传进来一个空的type,由系统去判定它是什么类型
	switch pc := usb.(type) {
	case PhoneConnect:
		pc.name = "我很好啊"
		fmt.Println("lxtSwitch :", pc.name)
	default:
		println("lxtSwitch")
	}
}
