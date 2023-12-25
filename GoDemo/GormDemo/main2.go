package main

import (
	"GormDemo/initDB"
	"gorm.io/gorm"
	"log"
)

type Dad struct {
	gorm.Model
	name string
	Son Son  `gorm:"foreignKey:id"`
}

type Son struct {
	gorm.Model
	SonName string

}

func main()  {
	var  u = new(Dad)
	initDB.DB.Preload("Son").First(&u).Debug()
	log.Println(u)
}



/*
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Class struct {
	gorm.Model
	ClassName string
	// 一对多 一个班级多个学生
	Students []Student
}

type Student struct {
	gorm.Model
	StudentName string
	// 一对多 学生根据班级id从属于哪个班级
	ClassId uint
	// 一对一 一个学生一个idcard
	IDCard IDCard
	// 多对多
	Teachers []Teacher `gorm:"many2many:student_teachers;"`
}
type IDCard struct {
	gorm.Model
	//一对一  一个id一个学生
	StudentID uint
	Num       int
}
type Teacher struct {
	gorm.Model
	// 多对多
	TeacherName string
	// 多对多
	Students []Student `gorm:"many2many:student_teachers;"`
}

func main() {
	// 数据库链接
	dsn := "root:1234@tcp(127.0.0.1:3306)/db_apiserver?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 自动迁移时，禁用外键约束
		SkipDefaultTransaction:                   false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 单数表名
			NoLowerCase:   false, // 不要小写转换
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(Student{}, Teacher{}, Class{}, IDCard{})
	db.Create(&Student{StudentName: "wang", ClassId: 1})
	var student Student
	db.Debug().First(&student)
	fmt.Println(student)
	var students []Student
	db.Debug().Where("ClassId < ?", 9).Find(&students)
	fmt.Println(students)
	db.Model(&student).

}
*/
