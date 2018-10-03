package handler

import (
	"net/http"
	"strconv"

	"github.com/gemcook/study-go/20181003/api_example/model"
	"github.com/gemcook/study-go/20181003/api_example/repository"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-xorm/xorm"
)

// PostGem 宝石登録ハンドラー
func PostGem(c *gin.Context) {
	// 前準備
	engine := c.MustGet("engine").(*xorm.Engine)
	repo := repository.NewGems(engine)

	// パラメータの読み取り
	var input model.Gem
	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		// 失敗応答
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// 登録処理
	created, err := repo.Create(&input)
	if err != nil {
		// 失敗応答
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// 成功応答
	c.JSON(http.StatusCreated, created)
}

// GetGems 宝石取得ハンドラー
func GetGems(c *gin.Context) {
	engine := c.MustGet("engine").(*xorm.Engine)
	repo := repository.NewGems(engine)
	list, err := repo.GetAll()
	if err != nil {
		// 失敗応答
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// 成功応答
	c.JSON(http.StatusOK, list)
}

// PutGem 宝石更新
func PutGem(c *gin.Context) {
	engine := c.MustGet("engine").(*xorm.Engine)
	repo := repository.NewGems(engine)

	// IDの読み取り
	id, err := strconv.ParseUint(c.Param("gem-id"), 10, 64)
	if err != nil {
		// 失敗応答
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// パラメータの読み取り
	var input model.Gem
	err = c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		// 失敗応答
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	updated, err := repo.Update(id, &input)
	if err != nil {
		// 失敗応答
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// 成功応答
	c.JSON(http.StatusOK, updated)
}
