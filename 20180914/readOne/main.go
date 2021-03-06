package main

import (
	"fmt"

	"github.com/gemcook/study-go/20180914/misc"
	"github.com/gemcook/study-go/20180914/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var gem models.Gem

	gem.ID = 1

	engine, err := misc.NewEngine()
	if err != nil {
		misc.ErrorMessage()
		panic(err)
	}

	_, err = engine.ID(gem.ID).Get(&gem)
	if err != nil {
		misc.ErrorMessage()
		panic(err)
	}

	fmt.Println(gem)
	misc.SuccessMessage()
}
