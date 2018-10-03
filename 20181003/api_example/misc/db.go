package misc

import (
	"fmt"
	"log"

	utils "github.com/gemcook/utils-go"
	"github.com/go-xorm/xorm"
	"github.com/joho/godotenv"
)

// NewEngine is connect DB
func NewEngine() (*xorm.Engine, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	engine, err := utils.InitMySQLEngine(utils.LoadMySQLConfigEnv())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return engine, nil
}
