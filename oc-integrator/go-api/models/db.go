package models

import (
	"zodi/api/core/common"

	gotools "github.com/boycehuang/go-tools"
	"github.com/boycehuang/go-tools/database"

	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

var RedisStore redis.Store
var MemStore memstore.Store

func ConnectRedis() {
	if gin.Mode() == gin.ReleaseMode {
		RedisStore, _ = redis.NewStore(10, "tcp", "10.0.0.117:6379", "", []byte(common.SESSION_SECRET))
	} else {
		MemStore = memstore.NewStore([]byte(common.SESSION_SECRET))
	}
}

func ConnectDB() {
	var dberr error

	if gin.Mode() == gin.ReleaseMode {
		DB, dberr = gotools.OpenDB(database.ConnectionConfig{
			Dialect: database.COCKROACHDB,
			Host:    "localhost",
			Port:    "26257",
			User:    "",
			Pass:    "",
			Dbname:  "",
		})
	} else {

		DB, dberr = gotools.OpenDB(database.ConnectionConfig{
			Dialect:  database.SQLITE,
			Filepath: "db.sqlite",
		})

		// DB, dberr = gotools.OpenDB(database.ConnectionConfig{
		// 	Dialect: database.POSTGRES,
		// 	Host:    "localhost",
		// 	Port:    "5432",
		// 	User:    "bbin",
		// 	Pass:    "1qaz@WSX",
		// 	Dbname:  "bbindb",
		// })
	}

	if dberr != nil {
		panic(dberr)
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(200)
	// DB.DropTable(&IrModel{}, &IrField{}, &IrAccess{})
	// DB.AutoMigrate(&IrModel{}, &IrField{}, &IrAccess{})

	// DB.DropTable(&Menu{}, &Site{}, &Role{}, &Layer{}, &GameHall{}, &Currency{}, &User{})
	// CreateTable(&Menu{}, &Site{}, &Role{}, &Layer{}, &GameHall{}, &Currency{}, &User{})
}
