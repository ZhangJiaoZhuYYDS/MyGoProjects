package main

import (
	"GormDemo/initDB"
	"gorm.io/gorm"
	"log"
)

// User 有多张 CreditCard，UserID 是外键
type Us struct {
	gorm.Model
	Credits []Credit
}

type Credit struct {
	gorm.Model
	Number string
	UsID uint
}


func main() {
	err := initDB.DB.AutoMigrate(&Us{},&Credit{})
	if err != nil {
		return 
	}

	//c1 := Credit{
	//	Number: "123456",
	//}
	//c2 := Credit{
	//	Number: "8791265",
	//}
	//u := Us{
	//	Credits: []Credit{c1,c2},
	//}
	//initDB.DB.Create(&u)
	var us Us
	var cd Credit
	//var a []Credit
	//initDB.DB.Debug().Model(&Us{}).Association("Credits").Find(&cd)
	//log.Println(us)


	initDB.DB.Debug().Preload("Credits").Find(&us)
	initDB.DB.Debug().Preload("Credits").Find(&cd)
	log.Println(us)
	log.Println(cd)
	//initDB.DB.Debug().Model(&Us{}).Preload("Credits").Find(&us)
	//initDB.DB.Debug().Model(&Us{}).Preload("Credit").Find(&cd)
	//log.Println(us)
	//log.Println(cd)
}
