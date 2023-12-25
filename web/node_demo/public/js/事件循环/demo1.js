// 引入 events 模块
var events = require('events');
// 创建 eventEmitter 对象
var eventEmitter = new events.EventEmitter();
// 创建事件处理程序
var eventHandler = function connected() {
    console.log('第一次调用：连接成功。');

    // 触发 data_received 事件 
    eventEmitter.emit('data_received');
}
// 绑定eventname事件事件及事件的处理程序
eventEmitter.on('eventName', eventHandler);
// 使用匿名函数绑定data_received事件
eventEmitter.on("data_received",function () {
    console.log(" 第二次调用：数据接收成功")
})

// 触发事件
eventEmitter.emit('eventName');