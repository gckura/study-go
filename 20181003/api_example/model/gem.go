package model

import "time"

// Gem is model struct
type Gem struct {
	ID            uint64    `xorm:"id pk autoincr" json:"id"`
	Name          string    `xorm:"name" json:"name"`
	JapaneseName  string    `xorm:"japanese_name" json:"japanese_name"`
	EnglishName   string    `xorm:"english_name" json:"english_name"`
	Hardness      float32   `xorm:"hardness" json:"hardness"`
	Price         int       `xorm:"price" json:"price"`
	ProducingArea string    `xorm:"producing_area" json:"producing_area"`
	Carat         float32   `xorm:"carat" json:"carat"`
	Weight        float32   `xorm:"weight" json:"weight"`
	Color         string    `xorm:"color" json:"color"`
	Memo          string    `xorm:"memo" json:"memo"`
	Scratched     *bool     `xorm:"scratched" json:"scratched"`
	MiningDate    time.Time `xorm:"mining_date updated" json:"mining_date"`
}

// TableName is function
func (Gem) TableName() string {
	return "t_gems"
}
