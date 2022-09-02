const http = require('http')

const server = http.createServer((req,res) => {
    res.writeHead(200, {'Content-Type' : 'text/html; charset=utf-8'})
    res.write('<h1>8080번 포트</h1>')
    res.end()
}).listen(8080)


const server2 = http.createServer((req, res) => {
    res.writeHead(200,{'Content-Type' : 'text/html; charset=utf-8'})
    res.write('<h1>8081번 포트</h1>')
}).listen(8081)


server.on('listening', () => {
    console.log('8080번 포트 실행')
})
server.on('error', (err) =>{
    console.log(err)
})

server2.on('listening', () => {
    console.log('8081번 포트 실행')
})

server2.on('error', (err) => {
    console.log(err)
})