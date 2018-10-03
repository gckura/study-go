package repository

import (
	"github.com/gemcook/study-go/20181003/api_example/model"
	"github.com/go-xorm/xorm"
)

type Gems struct {
	engine *xorm.Engine
}

func NewGems(engine *xorm.Engine) *Gems {
	g := Gems{
		engine: engine,
	}

	return &g
}

func (g *Gems) Create(input *model.Gem) error {
	_, err := g.engine.Insert(input)
	if err != nil {
		return err
	}

	return nil
}

func (g *Gems) GetAll() ([]*model.Gem, error) {
	var list []*model.Gem
	err := g.engine.Find(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (g *Gems) Delete(id uint64) error {
	var gem model.Gem

	_, err := g.engine.Where("id = ?", id).Delete(&gem)
	if err != nil {
		return err
	}
	return nil
}
