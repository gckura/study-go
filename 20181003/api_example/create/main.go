package main

import (
	"github.com/gemcook/study-go/20180914/misc"
	"github.com/gemcook/study-go/20180914/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	scratched := false
	gem := models.Gem{
		Name:          "サファイア",
		JapaneseName:  "サファイア",
		EnglishName:   "saphire",
		Hardness:      9,
		Price:         999,
		ProducingArea: "",
		Carat:         2.2,
		Weight:        22,
		Color:         "blue",
		Memo:          "Hello",
		Scratched:     &scratched,
	}

	engine, err := misc.NewEngine()
	if err != nil {
		misc.ErrorMessage()
		panic(err)
	}

	_, err = engine.Insert(&gem)
	if err != nil {
		misc.ErrorMessage()
		panic(err)
	}

	misc.SuccessMessage()
}
