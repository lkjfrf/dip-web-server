package db

import (
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB_Ins *DBManager
var DB_once sync.Once

type DBManager struct {
	DBMS *gorm.DB
}

func GetDBManager() *DBManager {
	DB_once.Do(func() {
		DB_Ins = &DBManager{}
	})
	return DB_Ins
}

func (db *DBManager) Init() {
	dsn := "root:Tomatosoup22!@tcp(127.0.0.1:3306)/dip?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "DIPADM:P!ssw0rd@tcp(10.5.147.148:3306)/dipdb?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			//LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true, // Ignore ErrRecordNotFound error for logger
			//Colorful:                  false,         // Disable color
		},
	)
	dbms, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		log.Println(err)
	} else {
		db.DBMS = dbms
		log.Println("DBConnect Success")
	}

	db.DBMS.AutoMigrate(&ZCMUSER{})
	db.DBMS.AutoMigrate(&SecurityPhrase{})
	db.DBMS.AutoMigrate(&WebLogin{})
}
