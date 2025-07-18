package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"log"
	"time" // 添加 time 包用于设置连接超时
)

var DB *gorm.DB

func Mysql(User, Password, Host, DB string, Port int) {
	// DSN (Data Source Name) 包含数据库连接所需的全部参数
	// 格式: user:password@protocol(address)/dbname?param=value
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", User, Password, Host, Port, DB)

	// sql.Open() 初始化数据库连接池，但不会立即建立实际连接
	// 返回的 *sql.DB 对象是协程安全的，可以全局使用
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("数据库连接初始化失败: %v", err)
	}
	defer db.Close() // 程序结束时关闭连接池，释放资源

	// 配置连接池参数（可选但推荐）
	db.SetMaxOpenConns(10)                  // 限制同时打开的最大连接数
	db.SetMaxIdleConns(5)                   // 限制空闲连接池中保留的最大连接数
	db.SetConnMaxLifetime(30 * time.Minute) // 设置连接的最大生命周期，避免长时间使用陈旧连接

	// 使用 Ping() 测试实际连接是否成功
	if err := db.Ping(); err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}
	fmt.Println("成功连接到 MySQL 数据库!")

	// 执行简单查询示例 - 获取数据库版本信息
	var version string
	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		log.Fatalf("查询执行失败: %v", err)
	}
	fmt.Printf("MySQL 版本: %s\n", version)

}
