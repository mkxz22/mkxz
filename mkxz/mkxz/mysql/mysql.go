package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time" // 添加 time 包用于设置连接超时
)

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
)

var (
	err error
	DB  *gorm.DB
)

func MysqlInit(user string, password string, host string, port int, database string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := DB.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return DB
}
