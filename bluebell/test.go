// @Author zhangjiaozhu 2023/6/4 15:05:00
package main

func main() {
	letterCombinations("23")
}

// 构造号码与字母映射
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// 结果集
var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	// 开始递归，传入参数，从第0个下标开始
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	// 递归到最后一个字符，递归结束
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		// 找到当前字符
		digit := string(digits[index])
		// 找到当前字符对应map映射
		letters := phoneMap[digit]
		// 统计映射长度
		lettersCount := len(letters)
		// 根据map映射长度递归调用几次
		for i := 0; i < lettersCount; i++ {
			// 传入当前字符下一个下标，并不断加上map映射字母，携带着传入下一个递归
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}
