package main

import (
	"github.com/gemcook/study-go/20180914/misc"
	"github.com/gemcook/study-go/20180914/models"
	_ "github.com/go-sql-driver/mysql"
)

// TableName is function
func (Gem) TableName() string {
	return "t_gems"
}

func main() {
	scratched := false
	gem := models.Gem{
		Name:          "Ruby",
		JapaneseName:  "ルビー",
		EnglishName:   "ruby",
		Hardness:      10,
		Price:         1200,
		ProducingArea: "",
		Carat:         1.2,
		Weight:        12,
		Color:         "red",
		Memo:          "Hei",
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
