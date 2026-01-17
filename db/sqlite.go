package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"super-send-tool/config/baseconfig"
)

var Db *gorm.DB
var err error

func SqLiteSetup(dbPath string) {
	_, err := os.Stat(dbPath)
	isFirst := false
	if err != nil {
		f, _ := os.Create(dbPath)
		defer f.Close()
		isFirst = true
	}
	conf := &gorm.Config{}
	if !baseconfig.CONFIG.Debug {
		conf.Logger = logger.Default.LogMode(logger.Silent)
	} else {
		conf.Logger = logger.Default.LogMode(logger.Info)
	}
	Db, err = gorm.Open(sqlite.Open(dbPath), conf)
	if isFirst {
		/*sqlTable := ``
		Db.Exec(sqlTable)*/
	}
	if err != nil {
		log.Fatalln(err)
	}
}
