package app

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"myblog/config"
	"myblog/pkg/database"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	var dbConfig gorm.Dialector
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		config.Viper.GetString("mysql.username"),
		config.Viper.GetString("mysql.password"),
		config.Viper.GetString("mysql.host"),
		config.Viper.GetString("mysql.port"),
		config.Viper.GetString("mysql.database"),
		config.Viper.GetString("mysql.charset"),
	)
	dbConfig = mysql.New(mysql.Config{
		DSN: dsn,
	})
	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))

	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(20)
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(300)
}
