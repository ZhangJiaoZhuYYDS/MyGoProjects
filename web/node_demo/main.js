route  =require('./route/view')
//console.log(route)
//path =require('path')
//console.log(path.resolve("./"));
//console.log(__filename)
const express = require('express')

const app = express()

app.use('/login',route);

app.use('/public',express.static('./views'))

app.listen(8888,()=>{
    console.log(8888)
})