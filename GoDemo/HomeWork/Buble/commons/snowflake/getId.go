// @Author zhangjiaozhu 2023/3/5 9:12:00
package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}

/*主程序中使用Init函数进行初始化

if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
	fmt.Printf("init sonwflake failed, err: %v\n", err)
	return
}
调用GenID()即可获得一个ID实例：
genID := GenID()
fmt.Printf("genID: %v\n", genID)
*/
