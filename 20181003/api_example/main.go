package main

import (
	"fmt"

	"github.com/gemcook/study-go/20181003/api_example/server"
)

func main() {
	err := server.Start()

	if err != nil {
		fmt.Println(err)
	}
}
