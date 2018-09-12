package main

import (
	"fmt"
	"time"

	"github.com/gemcook/study-go/20180914/utils"
	_ "github.com/go-sql-driver/mysql"
)

// Gem is model struct
type Gem struct {
	ID            int       `xorm:"id pk autoincr"`
	Name          string    `xorm:"name"`
	JapaneseName  string    `xorm:"japanese_name"`
	EnglishName   string    `xorm:"english_name"`
	Hardness      float64   `xorm:"hardness"`
	Price         int       `xorm:"price"`
	ProducingArea string    `xorm:"producing_area"`
	Carat         float64   `xorm:"carat"`
	Wieght        float64   `xorm:"carat"`
	Color         string    `xorm:"color"`
	Memo          string    `xorm:"memo"`
	Scratched     bool      `xorm:"scratched"`
	MiningDate    time.Time `xorm:"minig_date updated"`
}

// TableName is function
func (Gem) TableName() string {
	return "t_gems"
}

func main() {
	gem := Gem{
		Name:          "Ruby",
		JapaneseName:  "ルビー",
		EnglishName:   "ruby",
		Hardness:      10,
		Price:         1200,
		ProducingArea: "",
		Carat:         1.2,
		Wieght:        12,
		Color:         "red",
		Memo:          "Hei",
		Scratched:     false,
	}

	engine, err := utils.InitMySQLEngine(utils.LoadMySQLConfigEnv())
	if err != nil {
		fmt.Println(err)
	}

	// fuck, err := engine.Insert(&gem)

	fmt.Println(engine.Insert(&gem))
}
