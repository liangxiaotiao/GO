package areyouok

import (
	"fmt"
)

func imfinl() {
	fmt.Println("爱你哦")
	fmt.Println("")
	a, b := 4, 5
	if a == b {
		fmt.Println("都比较大?")
	} else if a > b {
		fmt.Println("瓶子比较大")
	} else if a < b {
		fmt.Println("嗨妹比较大")
	}
	//for
	for i := 0; i < 5; i++ {
		fmt.Println("你猜我是第几", i)
	}
	//break
	i := 0
	for {
		if i > 3 {
			break
		}
		fmt.Println("你是第几呀", i)
		i++
	}

	fmt.Println("Break掉咯")

}
