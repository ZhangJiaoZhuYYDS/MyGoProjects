const router = require('express').Router();
const fs = require('fs');
const { resolve } = require('path');
var s = 666
router
    .get('/index.html', (req, res) => {
        // 这里设置utt-8否则返回的buffer数据格式，会自动下载
        fs.readFile(resolve('./')+'/views/index2.html', 'utf8', (err, data) => {
            if (err) {
                console.log(err);
                return;
            }
            res.end(data.toString());
        });
    });
module.exports = {router,s};
