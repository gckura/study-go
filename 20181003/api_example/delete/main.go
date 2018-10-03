package main

import (
	"github.com/gemcook/study-go/20180914/misc"
	"github.com/gemcook/study-go/20180914/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var gem models.Gem

	gem.ID = 35

	engine, err := misc.NewEngine()
	if err != nil {
		misc.ErrorMessage()
		panic(err)
	}

	_, err = engine.ID(gem.ID).Delete(&gem)
	if err != nil {
		misc.ErrorMessage()
		panic(err)
	}

	misc.SuccessMessage()
}
