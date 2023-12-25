// @Author zhangjiaozhu 2023/8/8 8:57:00
package 深度优先

// 从一个起始的可访问的节点，获取节点中可访问哪些节点（每次访问过的节点都记录一下，防止重复访问），然后遍历可访问的节点继续递归遍历访问。最终比较访问过的节点数量是否等于所有节点数量

var (
	num int    //  访问了的房间的数量
	vis []bool // 标记房间是否被访问过
)

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	num = 0
	vis = make([]bool, n)
	dfs(rooms, 0)
	return num == n
}
func dfs(rooms [][]int, x int) {
	vis[x] = true // 每次访问未访问过的节点就更新访问状态
	num++         // 更新访问房间数量
	// 遍历节点中可访问的节点，递归访问 ， 只访问未访问过的节点
	for _, it := range rooms[x] {
		if !vis[it] {
			dfs(rooms, it)
		}
	}
}

/*  广度优先
func canVisitAllRooms(rooms [][]int) bool {
    n := len(rooms)
    num := 0
    vis := make([]bool, n)
    queue := []int{}
    vis[0] = true
    queue = append(queue, 0)
    for i := 0; i < len(queue); i++ {
        x := queue[i]
        num++
        for _, it := range rooms[x] {
            if !vis[it] {
                vis[it] = true
                queue = append(queue, it)
            }
        }
    }
    return num == n
}

*/
