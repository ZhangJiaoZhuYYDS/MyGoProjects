function calculate(NumberContent) {
    //获取+的位置
    let index = NumberContent.indexOf('+')
    //加号存在
    if (index > -1) {
        debugger
        return (
            calculate(NumberContent.substring(0, index)) +
            calculate(NumberContent.substring(index + 1))
        )
    }
    //当式子不再有加号时
    index = NumberContent.lastIndexOf('-');
    if (index > -1) {
        return (
            calculate(NumberContent.substring(0, index)) -
            calculate(NumberContent.substring(index + 1))
        )
    }
    //乘号
    index = NumberContent.lastIndexOf('*')
    if (index > -1) {
        return (
            calculate(NumberContent.substring(0, index)) *
            calculate(NumberContent.substring(index + 1))
        )
    }
    //除号
    index = NumberContent.lastIndexOf('/');
    if (index > -1) {
        return (
            calculate(NumberContent.substring(0, index)) /
            calculate(NumberContent.substring(index + 1))
        )
    }
    //将结果转换为数字
    return Number(NumberContent)
}
calculate("1*3+4+5")