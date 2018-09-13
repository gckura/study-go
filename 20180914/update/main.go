package main

import (
	"github.com/gemcook/study-go/20180914/misc"
	"github.com/gemcook/study-go/20180914/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	gem := models.Gem{
		Name:      "ほえー",
		Scratched: true,
	}

	engine, err := misc.NewEngine()
	if err != nil {
		misc.ErrorMessage()
		panic(err)
	}

	_, err = engine.ID(34).Update(&gem)
	if err != nil {
		misc.ErrorMessage()
		panic(err)
	}

	misc.SuccessMessage()
}
