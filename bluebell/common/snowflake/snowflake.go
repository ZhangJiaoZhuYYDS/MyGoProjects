// @Author zhangjiaozhu 2023/6/4 9:51:00
package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

func Init(machineId int64) (error error) {
	var st time.Time
	st, error = time.Parse("2006-01-02", "2023-01-01")
	if error != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, error = snowflake.NewNode(machineId)
	return
}
func GenId() int64 {
	return node.Generate().Int64()
}
