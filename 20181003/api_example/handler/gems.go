package handler

import (
	"github.com/gemcook/study-go/20181003/api_example/misc"
	"github.com/gemcook/study-go/20181003/api_example/model"
	"github.com/gemcook/study-go/20181003/api_example/repository"
	"github.com/gin-gonic/gin"
)

func PostGem(c *gin.Context) {
	engine, err := misc.NewEngine()
	if err != nil {
		c.JSON(500, "データベース接続異常")
		return
	}

	gemsRepo := repository.NewGems(engine)

	var input model.Gem

	err = c.ShouldBindJSON(&input)
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
