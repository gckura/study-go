package main

import (
	"github.com/gemcook/study-go/20180914/misc"
	"github.com/gemcook/study-go/20180914/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	scratched := true
	gem := models.Gem{
		ID:        37,
		Name:      "ほえー",
		Scratched: &scratched,
	}

	engine, err := misc.NewEngine()
	if err != nil {
		misc.ErrorMessage()
		panic(err)
	}

	_, err = engine.ID(gem.ID).Update(&gem)
	if err != nil {
		misc.ErrorMessage()
		panic(err)
	}

	misc.SuccessMessage()
}
