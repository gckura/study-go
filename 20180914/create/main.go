package main

import (
	"time"

	"github.com/gemcook/study-go/20180914/misc"
	"github.com/gemcook/study-go/20180914/models"
	_ "github.com/go-sql-driver/mysql"
)

// Gem is model struct
type Gem struct {
	ID            int       `xorm:"id pk autoincr"`
	Name          string    `xorm:"name"`
	JapaneseName  string    `xorm:"japanese_name"`
	EnglishName   string    `xorm:"english_name"`
	Hardness      float32   `xorm:"hardness"`
	Price         int       `xorm:"price"`
	ProducingArea string    `xorm:"producing_area"`
	Carat         float32   `xorm:"carat"`
	Weight        float32   `xorm:"weight"`
	Color         string    `xorm:"color"`
	Memo          string    `xorm:"memo"`
	Scratched     bool      `xorm:"scratched"`
	MiningDate    time.Time `xorm:"mining_date updated"`
}

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
