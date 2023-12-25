package main

/*给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
*/
func main() {
	s := "babad"
	println(m1(s))
}
// 1  暴力题解  遍历判断每一个子串
func m1(s string)  string{
	l :=len(s)
	// 边界判断
	if l < 2 {
		return s
	}
	// 暴力遍历每一个子串
	maxLen := 1  // 最长子串长度
	start := 0   // 记录每次子串起始位置
	for i:=0 ; i < l-1 ; i++ {
		for y := i+1 ; y < l ; y++ {
			if y - i + 1 > maxLen && validHuiWen(s[i:y+1]){
				maxLen=y-i+1
				start = i
			}
		}
	}
	return s[start:start+maxLen]

}
// 验证回文串
func validHuiWen(s string) bool {
	l := 0
	r := len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r++
	}
	return true
}
/*public class Solution {

    public String longestPalindrome(String s) {
        int len = s.length();
        if (len < 2) {
            return s;
        }

        int maxLen = 1;
        int begin = 0;
        // dp[i][j] 表示 s[i..j] 是否是回文串
        boolean[][] dp = new boolean[len][len];
        // 初始化：所有长度为 1 的子串都是回文串
        for (int i = 0; i < len; i++) {
            dp[i][i] = true;
        }

        char[] charArray = s.toCharArray();
        // 递推开始
        // 先枚举子串长度
        for (int L = 2; L <= len; L++) {
            // 枚举左边界，左边界的上限设置可以宽松一些
            for (int i = 0; i < len; i++) {
                // 由 L 和 i 可以确定右边界，即 j - i + 1 = L 得
                int j = L + i - 1;
                // 如果右边界越界，就可以退出当前循环
                if (j >= len) {
                    break;
                }

                if (charArray[i] != charArray[j]) {
                    dp[i][j] = false;
                } else {
                    if (j - i < 3) {
                        dp[i][j] = true;
                    } else {
                        dp[i][j] = dp[i + 1][j - 1];
                    }
                }

                // 只要 dp[i][L] == true 成立，就表示子串 s[i..L] 是回文，此时记录回文长度和起始位置
                if (dp[i][j] && j - i + 1 > maxLen) {
                    maxLen = j - i + 1;
                    begin = i;
                }
            }
        }
        return s.substring(begin, begin + maxLen);
    }
}
*/
// 2 中心扩展算法

// java
//class Solution {
///*1 中心扩散法*/
//public String longestPalindrome(String s) {
//// 判断为空
//if (s.length()<1) return "";
//int start = 0; // 起始位置
//int end = 0;  // 结束位置
//int len = 1;
//int MaxStart = 0; // 最长子串开始位置
//int MaxLen = 0; // 最长子串的长度
//int strLen = s.length();
//for ( int i = 0 ; i < strLen ; i++) {
//start = i-1; // 当前遍历元素的前一位
//end = i+1; // 当前遍历元素的后一位
//// 向左扩散与当当前元素相同的元素数，计算长度
//while ( start >= 0 && s.charAt(i) == s.charAt(start)){
//start--;
//len++;
//}
//// 向右扩散与当前元素相同的元素，计算个数
//while( end < strLen && s.charAt(i) == s.charAt(end)){
//end++;
//len++;
//}
//// 左右扩散寻找左右相同的回文字符
//while( start >= 0 && end < strLen && s.charAt(start) == s.charAt(end)){
//start--;
//end++;
//len+=2;
//}
//// 每次判断更新最大值
//if (len > MaxLen) {
//MaxLen = len;
//MaxStart = start;
//}
//// 每次结束刷新长度
//len = 1;
//}
//return s.substring(MaxStart+1,MaxStart+MaxLen+1);
//}
//}

// 3 动态规划
/*中心扩散的方法，其实做了很多重复计算。动态规划就是为了减少重复计算的问题。动态规划听起来很高大上。其实说白了就是空间换时间，将计算结果暂存起来，避免重复计算。作用和工程中用 redis 做缓存有异曲同工之妙。
我们用一个 boolean dp[l][r] 表示字符串从 i 到 j 这段是否为回文。试想如果 dp[l][r]=true，我们要判断 dp[l-1][r+1] 是否为回文。只需要判断字符串在(l-1)和（r+1)两个位置是否为相同的字符，是不是减少了很多重复计算。
进入正题，动态规划关键是找到初始状态和状态转移方程。
初始状态，l=r 时，此时 dp[l][r]=true。
状态转移方程，dp[l][r]=true 并且(l-1)和（r+1)两个位置为相同的字符，此时 dp[l-1][r+1]=true。
*/
public String longestPalindrome(String s) {
	// 边界判断
	if (s.length()<2){
		return "";
}
	// 动态规划
	int strLen = s.length();
	int Maxstart = 0; // 最长回文串的起点
	int Maxend = 0; // 最长回文串的终点
	int MaxLen = 0; // 最长回文串的长度
	// 新建dp数组，存储从一个到另一个元素位置是否为回文，若为回文就修改为true;
	boolean[][] dp = new boolean[strLen][strLen];
	for(int i = 1 ; i < strLen ; i++){
		for (int j = 0 ;j < i ; j++){
			if (s.charAt(i)== s.charAt(j) && (j-i<=2 || dp[i-1][j+1])){
				dp[i][j] = true;
				MaxLen = j - l +1;
				Maxstart = i;
				Maxend = j ;
}
}
}
	return s.substring(Maxstart,Maxend+1);
}