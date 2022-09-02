const http = require('http')
const fs = require('fs').promises

http.createServer(async (req,res) => {
    const data = await fs.readFile('./server2.html')
    res.writeHead(200, {'Content-Type' : 'text/html; charset=utf-8'})
    res.end(data)
    
}).listen(8081, () => {
    console.log('8081번 포트 실행')
})