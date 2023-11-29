package mysql

import (
	"fmt"
	"go-edu/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 全局数据库 db
var db *gorm.DB

// / 包初始化函数，可以用来初始化 gorm
func Init(cfg *config.MySQLConfig) (err error) {

	// 配置 dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 额外的连接配置
	sqlDB, err := db.DB() // database/sql.DB
	if err != nil {

		return
	}

	// 以下配置要配合 my.conf 进行配置
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return
	//sqlDB.SetConnMaxLifetime(time.Second * 5) //设置空闲断开时间
	//return

}

// 获取 gorm db，其他包调用此方法即可拿到 db
// 无需担心不同协程并发时使用这个 db 对象会公用一个连接，因为 db 在调用其方法时候会从数据库连接池获取新的连接
func GetDB() *gorm.DB {
	return db
}
