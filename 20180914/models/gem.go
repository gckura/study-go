package models

import "time"

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
	Wieght        float32   `xorm:"weight"`
	Color         string    `xorm:"color"`
	Memo          string    `xorm:"memo"`
	Scratched     bool      `xorm:"scratched"`
	MiningDate    time.Time `xorm:"mining_date updated"`
}

// TableName is function
func (Gem) TableName() string {
	return "t_gems"
}
