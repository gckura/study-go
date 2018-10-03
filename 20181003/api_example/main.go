package main

import (
	"fmt"
	"os"

	"github.com/gemcook/study-go/20181003/api_example/server"
)

func main() {
	err := server.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
