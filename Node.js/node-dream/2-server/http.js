const http = require('http')
const fs = require('fs')


const server = http.createServer((req,  res) => {
    console.log('서버 생성!!')
    console.log('요청 타입',req.method)
    console.log('요청 경로',req.url)

    const URL = req.url

    res.setHeader('Content-Type', 'text/html')


    if(URL === '/'){
        fs.createReadStream('./html/home.html').pipe(res)

        // res.write('<html>')
        // res.write('<h1>Welcome! Home</h1>')
        // res.write('<html>')
    }

    if(URL === '/courses'){
        res.write('<html>')
        res.write('<h1>Welcome! Courses</h1>')
        res.write('<html>')
    }



    res.end()
})




server.listen(8888)