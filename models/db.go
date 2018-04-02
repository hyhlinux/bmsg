package models

import (
	_ "github.com/lib/pq"
	"github.com/astaxie/beego/orm"
	"bmsg/logger"
	"bmsg/config"
)

func init()  {
	InitDB()
}


func InitDB() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		panic(err)
	}

	pgDataSource := config.AppConf.PgDataSource
	if pgDataSource == ""{
		pgDataSource = "user=postgres password=postgres dbname=test host=127.0.0.1 port=5432 sslmode=disable"
	}
	err = orm.RegisterDataBase("default", "postgres", pgDataSource)
	if err != nil {
		logger.Errorf("pgDataSource:%v", pgDataSource)
		panic(err)
	}
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
	ormDebug := false
	if config.AppConf.OrmDebug {
		ormDebug = true
	}
	orm.Debug = ormDebug
	logger.Debugf("ormDebug:%v pgDataUrl:%v", ormDebug, pgDataSource)
}
