package demo1

/*
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter , r *http.Request)  {
	// 直接以html格式返回
	//_, _ = fmt.Fprintln(w, "<h1>fefsef</h1>")
	// 读取文件形式返回
	file, _ := ioutil.ReadFile("./index.html")
	fmt.Println(file)
	fmt.Fprintln(w,string(file))

}
func main() {
	http.HandleFunc("/hello",sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http server failure,%v",err)
		return
	}
}
*/