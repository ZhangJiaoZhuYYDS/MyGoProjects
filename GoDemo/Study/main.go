package main

import (
	"Study/config/mysql"
	"Study/models"
	"Study/router"
	"log"
)


// @title 标题
// @version 1.0
// @description 描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	// 初始化数据库
	mysql.DB.AutoMigrate(&models.ProblemBasic{},&models.ProblemCategory{},&models.UserBasic{})

	var user models.ProblemCategory
	mysql.DB.Debug().Model(&models.ProblemBasic{}).Select("problem_category.problem_id, problem_category.category_id").Joins("right join problem_category on problem_category.category_id = problem_basic.id").Scan(&user)
	log.Println(user)
	//database := mysql.DB.Migrator().CurrentDatabase()
	//log.Println(database)
	//data := make([]models.ProblemBasic,0)
	//mysql.DB.Create(&models.ProblemBasic{
	//	Identity:   "problem_1",
	//	Title:      "文章标题",
	//	Content:    "正文",
	//	MaxRuntime: 3000,
	//	MaxMem:     3000,
	//})
	//
	//err2 := mysql.DB.Debug().Find(&data).Error
	//if err2 != nil {
	//	log.Println("查询数据失败")
	//}
	//for _ , v := range data {
	//	log.Println(v)
	//}
	//type Te struct {
	//	gorm.Model
	//	Identity string
	//	Name string
	//	ParentId int
	//	category *models.CategoryBasic
	//}


	r := router.Router()

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
