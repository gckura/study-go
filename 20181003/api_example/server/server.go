package server

import (
	"github.com/gemcook/study-go/20181003/api_example/handler"
	"github.com/gin-gonic/gin"
)

func Start() error {
	router := gin.Default()

	router.POST("/gems", handler.PostGem)

	err := router.Run(":8080")
	return err
}
