package main

// offer-05------替换空格
func main() {
	//输入：
	s := "We are happy."
	// 输出："We%20are%20happy."
	println(replaceSpace2(s))
}

// 1 新建字符串
func replaceSpace(s string) string {
	// 字符串转字节数组
	b := []byte(s)
	// 声明一个结果数组
	result := make([]byte, 0)
	// 遍历字节数组元素
	for i := range b {
		// 如果元素等于‘ ’ ,就把 字符串转为[]byte{'%','2','0'} 利用go语法“...” 把数组内元素依次添加到结果数组内
		if s[i] == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, b[i])
		}
	}
	return string(result)
}

// 2 双指针原地修改
// 2.1遍历原数组b，统计数组内空格的数量
// 2.2拓宽原数组，增加新长度（长度为：原数组长度+空格数量*2）
// 2.3设置双指针，l起始位置为原数组的最后一位，r为数组n的最后一位
// 2.4遍历数组：若元素不为空，就把l指针元素添加到数组，然后l,r同时左移减一。若元素为空，就把’%‘‘2’‘0’添加到数组，然后l左移一位，r左移动三位
func replaceSpace2(s string) string {
	length := len(s)
	b := []byte(s)
	spaceCount := 0
	// 2.1
	for i := range b {
		if b[i] == ' ' {
			spaceCount++
		}
	}
	// 2.2
	temp := make([]byte, spaceCount*2)
	b = append(b, temp...)
	// 2.3
	l := length - 1
	r := len(b) - 1
	// 2.4
	for l >= 0 {
		if b[l] != ' ' {
			b[r] = b[l]
			l--
			r--
		} else {
			b[r] = '0'
			r--
			b[r] = '2'
			r--
			b[r] = '%'
			r--
			l--
		}
	}
	return string(b)
}
