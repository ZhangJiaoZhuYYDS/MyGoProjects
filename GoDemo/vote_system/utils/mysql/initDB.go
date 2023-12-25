// @Author zhangjiaozhu 2023/3/18 13:57:00
package mysql

import (
	"github.com/lexkong/log"
	viper2 "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"vote_system/models"
)

var DB *gorm.DB

func InitDB()  error{
	root := viper2.GetString("mysql.root")
	password := viper2.GetString("mysql.password")
	mydb := viper2.GetString("mysql.db")
	dsn := root+":"+password+"@tcp(127.0.0.1:3306)/"+mydb+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	DB = db
	err = DB.AutoMigrate(&models.User{},&models.Votes{},&models.Votesopt{})
	if err != nil {
		log.Info("数据库建表失败")
		return err
	}
	return nil
}
