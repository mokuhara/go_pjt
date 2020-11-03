package repository

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"goPjt/model"
	"log"
	"os"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "postgres"
	err := godotenv.Load("env/dev.env")
	if err != nil {
		log.Fatal(err.Error())
	}
	DsName := os.Getenv("DSN")
	err = errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != ""{
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.User))
	DbEngine.Sync2(new(model.Profile))
	fmt.Println("init database success!")
}

