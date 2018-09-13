package main

import (
	"fmt"

	"github.com/gemcook/study-go/20180914/misc"
	"github.com/gemcook/study-go/20180914/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var gems []models.Gem

	engine, err := misc.NewEngine()
	if err != nil {
		misc.ErrorMessage()
	}

	err = engine.Find(&gems)
	if err != nil {
		misc.ErrorMessage()
	}

	fmt.Println(gems)
	misc.SuccessMessage()
}
