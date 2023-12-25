package svc
// 上下文 加载配置文件;加载模型方法
import (
	"book/service/user/api/internal/config"
	"book/service/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	// 配置文件
	Config config.Config
	// 模型方法
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 获取数据库连接，redis连接。然后加载到上下文中
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(conn,c.CacheRedis),
	}
}
