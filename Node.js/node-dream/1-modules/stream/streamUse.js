const http = require('http')
const fs = require('fs')


const server = http.createServer((req, res) => {
    const stream = fs.createReadStream('./file.txt')
    stream.pipe(res)
})

server.listen(3000)