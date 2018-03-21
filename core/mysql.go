package core

import (
	"math/rand"
	"gogrpcgin/conf"
	"gogrpcgin/utils"

	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

func MasterDB(dbName string) (dbResource *xorm.Engine){

	if dbName == "" {
		dbName = "comic"
	}

	dbResource, err := xorm.NewEngine("mysql", conf.Conf.DB[dbName]["master"][rand.Intn(len(conf.Conf.DB[dbName]["master"]))])
	utils.ErrToPanic(err)

	return dbResource
}


func SlaveDB(dbName string) (dbResource *xorm.Engine){

	if dbName == "" {
		dbName = "comic"
	}

	dbResource, err := xorm.NewEngine("mysql", conf.Conf.DB[dbName]["slave"][rand.Intn(len(conf.Conf.DB[dbName]["slave"]))])
	utils.ErrToPanic(err)

	return dbResource
}