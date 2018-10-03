package server

import (
	"github.com/gemcook/study-go/20181003/api_example/handler"
	"github.com/gemcook/study-go/20181003/api_example/misc"
	"github.com/gin-gonic/gin"
)

// Start APIサーバー開始
func Start() error {

	router := gin.Default()
	r := router.Group("/", setEngineMiddleware())

	r.POST("/gems", handler.PostGem)
	r.GET("/gems", handler.GetGems)
	r.PUT("/gems/:gem-id", handler.PutGem)

	err := router.Run(":8080")
	return err
}

func setEngineMiddleware() gin.HandlerFunc {
	engine, err := misc.NewEngine()
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		c.Set("engine", engine)
		c.Next()
	}
}
