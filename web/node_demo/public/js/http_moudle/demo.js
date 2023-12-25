// 引入http模块
const http = require("http");
// 创建服务器实例
var server = http.createServer();

//使用服务器实例的.on()方法，为服务器绑定一个request事件
server.on('request',(req,res)=>{
//只要有客户端来请求我们自己的服务器，就会触发request事件，从而调用这个事件处理函数
//req是请求对象，它包含了与客户端相关的数据和属性，例如：
//req.url是客户端相关的数据和属性，例如：
//req.url是客户端请求的url地址
//req.method 是客户端的method请求类型
    console.log('someone visit our web server.')
    const str=`Your request url is ${req.url},and request method is ${req.method},req是${req}`
    console.log(str)
    //为了防止中文显示乱码的问题，需要设置响应头
    res.setHeader('Content-Type','text/html;charset=utf-8')
    //向客户端发送指定内容，并结束这次请求的处理过程
    res.end(str)
})

//调用server.listen(端口号，cb回调）方法，即可启动web服务器
server.listen(80,()=>{
    console.log('http server running at http://127.0.0.1')
})



