package server

import (
	"github.com/gemcook/study-go/20181003/api_example/handler"
	"github.com/gemcook/study-go/20181003/api_example/misc"
	"github.com/gin-gonic/gin"
)

func Start() error {
	router := gin.Default()

	r := router.Group("/", EngineMiddleware())

	r.POST("/gems", handler.PostGem)
	r.GET("/gems", handler.GetGems)
	r.DELETE("/gems/:gem-id", handler.DeleteGem)

	err := router.Run(":8080")
	return err
}

func EngineMiddleware() gin.HandlerFunc {
	engine, err := misc.NewEngine()
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		c.Set("engine", engine)
		c.Next()
	}
}
