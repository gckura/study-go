package handler

import (
	"fmt"
	"strconv"

	"github.com/gemcook/study-go/20181003/api_example/model"
	"github.com/gemcook/study-go/20181003/api_example/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

func PostGem(c *gin.Context) {
	engine := c.MustGet("engine").(*xorm.Engine)
	gemsRepo := repository.NewGems(engine)

	var input model.Gem

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, "リクエストボディのBind失敗: "+err.Error())
		return
	}

	err = gemsRepo.Create(&input)
	if err != nil {
		c.JSON(400, "登録失敗: "+err.Error())
		return
	}

	c.JSON(201, "登録成功")
}
func GetGems(c *gin.Context) {
	engine := c.MustGet("engine").(*xorm.Engine)
	gemsRepo := repository.NewGems(engine)

	list, err := gemsRepo.GetAll()
	if err != nil {
		c.JSON(400, "取得失敗: "+err.Error())
		return
	}

	c.JSON(200, list)
}

func DeleteGem(c *gin.Context) {
	engine := c.MustGet("engine").(*xorm.Engine)

	gemsRepo := repository.NewGems(engine)

	idStr := c.Param("gem-id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, "id が数値じゃない")
		return
	}
	err = gemsRepo.Delete(id)
	if err != nil {
		c.JSON(400, fmt.Sprintf("id = %vの削除失敗", id))
		return
	}

	c.JSON(204, nil)
}
