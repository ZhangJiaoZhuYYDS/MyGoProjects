package util

import "testing"

// 测试文件  测试文件以_test.go结尾，和待测试源码放在同一文件夹下
/*·文件名必须是test.go结尾，跟源文件在同一个包。
·测试用例函数必须以Test、Benchmark、Example开头
·执行测试用例时的顺序，会按照源码中的顺序依次执行
·单元测试函数TestXxx0的参数是testing.T,可以使用该类型来记录错误或测试状态
·性能测试函数BenchmarkXxx0的参数是testing.B,函数内以b.N作为循环次数，其中N会动态变化
·示例函数ExampleXxx0没有参数，执行完会将输出与注释Output:进行对比
·测试函数原型：func TestXxx(t*testing.T),Xx部分为任意字母数字组合，首字母大写，例如：
TestgenShortId是错误的函数名，TestGenShortId是正确的函数名
·通过调用testing.T的Error、Errorf、.FailNow、Fatal、Fatallf方法来说明测试不通过，通过调用Log、Logf
方法来记录测试信息：*/
func TestGenShortId(t *testing.T) {
	shortid, err := GenShortId()
	if shortid == "" || err != nil {
		t.Error("failed")
	}
	t.Log("pass")
}
