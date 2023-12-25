package initDB

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
//type User struct {
//	ID           uint
//	Name         string
//	Email        *string
//	Age          uint8
//	Birthday     time.Time
//	MemberNumber sql.NullString
//	ActivatedAt  sql.NullTime
//	CreatedAt    time.Time
//	UpdatedAt    time.Time
//}

var DB *gorm.DB
//
//type Us struct {
//	gorm.Model
//	Name      string
//	CompanyID string
//	Company   Company `gorm:"references:Code"` // 使用 Code 作为引用
//}

//type Company struct {
//	ID   int
//	Code string
//	Name string
//}


// `User` 属于 `Company`，`CompanyID` 是外键
//type User struct {
//	gorm.Model
//	Name      string
//	CompanyID int
//	Company   Company
//}
//
//type Company struct {
//	ID   int
//	Name string
//}
// User 有一张 CreditCard，UserID 是外键
type User struct {
	gorm.Model
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}


func init() {
	// 数据库链接
	dsn := "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表名
			NoLowerCase:   false, // 不要小写转换
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)

	DB = db // 将db赋值给全局DB

	// 迁移 schema 结构体字段和数据库表数据的绑定
	//err = db.AutoMigrate(&Product{})
	//if err != nil {
	//	fmt.Println("数据库创建失败")
	//}
	//
	//DB.Debug().AutoMigrate(&User{})
	//DB.Debug().AutoMigrate(&User{},&CreditCard{})
	//
	//c := CreditCard{
	//	Number: "123456",
	//}
	//
	//u := User{
	//	CreditCard: c,
	//}
	//db.Create(&u)

	//user := User{
	//	Name:            "jinzhu",
	//	CompanyID: 1,
	//	Company: Company{
	//		ID:   0,
	//		Name: "666",
	//	},
	//}

	//DB.Debug().Create(&user)
	//DB.Debug().Create(&User{Name: "wang",Company: Company{
	//	ID:   0,
	//	Name: "li",
	//},
	//})

	//var u User
	//DB.Debug().Model(&User{}).First(&u)
	//fmt.Println(u)
	//DB.Debug().Model(&User{}).Preload("Company").First(&u)
	//fmt.Println(u)
	// BEGIN TRANSACTION;
	// INSERT INTO "addresses" (address1) VALUES ("Billing Address - Address 1"), ("Shipping Address - Address 1") ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "users" (name,billing_address_id,shipping_address_id) VALUES ("jinzhu", 1, 2);
	// INSERT INTO "emails" (user_id,email) VALUES (111, "jinzhu@example.com"), (111, "jinzhu-2@example.com") ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "languages" ("name") VALUES ('ZH'), ('EN') ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "user_languages" ("user_id","language_id") VALUES (111, 1), (111, 2) ON DUPLICATE KEY DO NOTHING;
	// COMMIT;


















	// 1 Create  建表
	//err = DB.Debug().Create(&Product{Code: "D4", Price: 10}).Error
	//if err != nil {
	//	fmt.Println("建表失败")
	//} else {
	//	fmt.Println("建表成功")
	//}
	//
	//DB.Debug().AutoMigrate(&Us{})
	//DB.Debug().Create(&Us{
	//	Name:      "wang",
	//	CompanyID: "123",
	//	Company:   Company{},
	//})

	// 建表时只更新部分字段
	//DB.Select("Price").Debug().Create(&Product{
	//	Code:  "12",
	//	Price: 0,
	//})
	// 建表时不更新某些字段
	//DB.Omit("Price").Debug().Create(&Product{
	//	Code:  "12",
	//	Price: 0,
	//})

	// 2 Read  查表
	//var product Product
	//db.First(&User{}, 3) // 根据整型主键查找
	//db.First(&Product{}, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//db.First(&product)
	//fmt.Println(product)
	//db.Take(&Product{})

	// 3 Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 300)

	// 4  Delete - 删除 product
	// db.Delete(&product, 1)

	//var user User
	//db.First(&user, 3)
	//fmt.Println(user)
	// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	// 创建记录并更新给出的字段。
	//tx := db.Create(&user) // 通过数据的指针来创建
	//fmt.Println(tx == db)
	//db.Select("Name", "Age", "CreatedAt").Create(&user)
	// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")

}
