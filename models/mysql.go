package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitMySQL() (err error) {
	dsn := "root:qwe123456@tcp(127.0.0.1:3307)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 Mariadb 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 Mariadb 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		//DriverName: "my_mysql_driver", // 自定义mysql驱动
	}), &gorm.Config{})
	//DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	// 获取通用数据库对象 sql.db ，然后使用其提供的功能
	sqlDB, err := DB.DB()
	if err != nil {
		return
	}else {
		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)
		return sqlDB.Ping()
	}
}