const express = require('express')
const router = express.Router()

router.get('/', (req, res) => {
    res.render('index') //index.html 파일을 렌더링 하여 화면에 보여주겠다는 문법
})


module.exports = router