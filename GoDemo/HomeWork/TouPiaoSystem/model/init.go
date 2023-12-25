package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

// 初始化获取数据库连接
func (db *Database) Init() {
	DB = &Database{Self: InitSelfDB()}
	err := DB.Self.AutoMigrate(&Votes{},&Votesopt{},&VotesVotesopt{})
	if err != nil {
		return 
	}
}

// 读取数据库配置文件
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

// 连接数据库
func openDB(username, password, addr, name string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		// 时区
		"Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "",    // 表名前缀
			SingularTable: true,  // 单数表名
			//NameReplacer:  nil,   // 名字替换
			//NoLowerCase:   false, // 大小写转换
		},
	})
	if err != nil {
		fmt.Println(err, "数据库链接失败")
	}
	// setupDB(db)
	return db
}
