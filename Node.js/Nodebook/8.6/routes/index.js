// [ GET / 로 접속했을때의 라우터 ]

const express = require('express')
const User = require('../schemas/user.js') 

const router = express.Router()

router.get('/', async(req, res, next) => {
    try{
        //User 스키마에서 데이터를 다 찾아온다.
        const users = await User.find({})
        // mongoose.html에 가져온 users데이터를 변수로 넣는다.
        res.render('mongoose', { users })
    }
    catch(err){
        console.log(err)
        next(err)
    }
})

module.exports = router;