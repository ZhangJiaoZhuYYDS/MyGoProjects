// @Author zhangjiaozhu 2023/6/4 8:41:00
package models

var s string = "123"

func main() {
	test("123")
}

func test(s string) {
	mp := map[string]string{
		"1": "abc",
		"2": "def",
		"3": "qwr",
	}
	var r = func(string2 string, index int) {}
	r = func(string2 string, index int) {
		for i := 0; i < len(s); i++ {
			c := mp[string(s[i])]
			cl := len(c)
			for j := 0; j < cl; j++ {
				string2 += string(c[j])
				r(string2, index+1)
			}
		}
	}
	r("", 0)

}
