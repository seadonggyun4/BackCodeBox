const path = require('path')
const express = require('express')//내부에 http 모듈이 내장되어있음
const app = express()


app
.set('port', process.env.PORT || 3000)
.get('/', (req,res) => {
    //res.send('<h1>Hello Express</h1>')
    res.sendFile(path.join(__dirname, '/index.html'))
})
.get('/about', (req,res) => {
    res.sendFile(path.join(__dirname, '/about.html'))
})
.listen(app.get('port'), () => {
    console.log(app.get('port'), '번 포트에서 대기중')
})


/*
app.set(키,값) => app.get으로 가져올수 있는데이터 저장 

app.get(주소, 라우터) =>  해당주소에 접근했을때 수행할 동작/ express 에서는 res.write, res.end 대신 res.send를 사용

express 의 라우터 => post, put, patch, delete, option 에 대한 동작을 app.post, app.put, app.patch, app.delete, app.option

res.sendFile => 단순 문자열 대신 HTML과 같은 파일로 응답하고 싶을때/ path 모듈을 사용해 경로 지정/ send 와 sendFile 을 같이사용하면 안된다. 

*/