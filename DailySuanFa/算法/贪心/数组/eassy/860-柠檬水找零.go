// @Author zhangjiaozhu 2023/8/15 10:38:00
package eassy

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0 // 动态维护5元和10元的数量

	for _, bill := range bills {
		// 1 如果账单是5，直接收下
		if bill == 5 {
			five++
			// 2 如果账单是10 先判断5的数量是否满足，满足就5-- 10++
		} else if bill == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
			// 3 账单是20 就分情况讨论，优先给10 5 再其次给 5 5 5
		} else {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}