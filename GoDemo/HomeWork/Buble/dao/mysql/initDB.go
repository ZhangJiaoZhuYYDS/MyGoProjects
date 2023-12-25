// @Author zhangjiaozhu 2023/3/4 21:21:00
package mysql

// 使用config.GetDB()获取数据库连接池对象
import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 定义一个全局的gorm模型
var _db *gorm.DB

// 定义全局日志输出
var ormLogger logger.Interface

func init() {
	// 日志级别
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	// 声明err变量，下面不能使用:=，否侧_d变量会被当做局部变量使用，导致外部无法访问_db变量
	var err error
	// 数据库链接
	dsn := "root:1234@tcp(127.0.0.1:3306)/db_apiserver?charset=utf8mb4&parseTime=True&loc=Local"
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 ormLogger, // 日志级别
		SkipDefaultTransaction: false,
		NamingStrategy:         schema.NamingStrategy{
			//SingularTable: true, // 单数表名
			//NoLowerCase:   true, // 表名不要下划线
		},
	})
	if err != nil {
		panic("连接数据库失败，error :" + err.Error())
	}
	// 设置连接池信息
	sqlDB, _ := _db.DB()       // 获取连接池 1） 针对大量用户并发操作数据库出现连接超时问题：2）针对少量用户快速实现相关数据库操作的问题：一般情况上面的两个函数是一起使用的，而且最大连接数的设置，必须要大于最大可空闲连接数。
	sqlDB.SetMaxOpenConns(100) // 1)最大数量 SetMaxOpenConns（）表示最大的连接数，这个我们不设置默认就是不限制，可以无限创建连接，问题就在数据库本身有瓶颈，无限创建，会损耗性能。所以我们要根据我们自己的数据库瓶颈情况来进行相关的设置。当出现连接数超出了我们设定的数量时候，后面的用户等待超时时间之前，有连接释放就会自动获得操作的权限，否则返回连接超时。（每个公司的使用情况不同，所以根据情况自己设定，个人建议不要采用默认无限制创建连接）
	sqlDB.SetMaxIdleConns(20)  // 2) 空闲连接池对象数量 SetMaxIdleConns（）表示设置最大的可空闲连接数，该函数的作用就是保持等待连接操作状态的连接数，这个主要就是避免操作过程中频繁的获取连接，释放连接。默认情况下会保持的连接数量为2.就是说会有两个连接一直保持，不释放，等待需要使用的用户使用。

}

// 获取gorm db 对象 ，其他包需要执行数据库查询的时候，只需要包名调用getDB()方法获取db对象即可
// 不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetDB() *gorm.DB {
	return _db
}
