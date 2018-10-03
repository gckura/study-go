package repository

import (
	"fmt"

	"github.com/gemcook/study-go/20181003/api_example/model"
	"github.com/go-xorm/xorm"
)

// Gems 宝石リポジトリー
type Gems struct {
	engine *xorm.Engine
}

// NewGems 宝石リポジトリ初期化
func NewGems(engine *xorm.Engine) *Gems {
	return &Gems{
		engine: engine,
	}
}

// Create 宝石登録
func (g *Gems) Create(input *model.Gem) (created *model.Gem, err error) {
	_, err = g.engine.Insert(input)
	if err != nil {
		return nil, err
	}

	return input, nil
}

// GetAll 宝石一覧取得
func (g *Gems) GetAll() ([]*model.Gem, error) {
	var list []*model.Gem
	err := g.engine.Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// Update 宝石更新
func (g *Gems) Update(id uint64, input *model.Gem) (*model.Gem, error) {
	fmt.Printf("%v %+v", id, *input)
	_, err := g.engine.Where("id = ?", id).Update(input)
	if err != nil {
		return nil, err
	}

	var updated model.Gem
	_, err = g.engine.Where("id = ?", id).Get(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}
