package main

import (
	"fmt"
)

func main() {
	msg := "Hello world"
	for i := 0; i < 5; i++ {
		if i == 3 {
			fmt.Println("さーん！！！")
		} else {
			fmt.Println(msg, i)
		}
	}
}
