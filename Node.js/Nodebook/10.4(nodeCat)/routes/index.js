const express = require('express')
const axios = require('axios')

const router = express.Router()


// [ 토큰 테스트 라우터 ]
// 데이터를 받아올 것이기때문에 비동기 통신 
router.get('/test', async (req, res, next) => {
    try{
        // 세션에 토큰이 없으면 발급 시도
        if(!req.session.jwt){
            const tokenResult = await axios.post('http://localhost:8002/v1/token', {
                clientSecret: process.env.CLIENT_SECRET,
            })
            // 토큰 발급에 성공하면
            if(tokenResult.data && tokenResult.data.code === 200){
                req.session.jwt = tokenResult.data.token // 세션에 토큰 저장
            } else { // 토큰 발급에 실패하면
                return res.json(tokenResult.data) //발급 실패 사유 응답
            }
        }
        // 결과값 받아오기
        const result = await axios.get('http://localhost:8002/v1/test', {
            headers: { authorization: req.session.jwt},
        })

        return res.json(result.data) // 결과값 표출
    } catch (error){
        console.error(error)

        if(error.response.status === 419){ // 토큰만료시
            return res.json(error.response.data) //에러메시지 본문 표출
        }

        return next(error)
    }
})


module.exports = router