// [ 서버 ]

const express = require('express')
const http = require('http') // socket.io는 웹소켓 이기 때문에 http기반으로 통신되어야 한다.
const app = express()
const path = require('path')
const server = http.createServer(app) //app은 http 통신 기반으로 이루어진다.

const socketIO = require('socket.io')
const io = socketIO(server) // socketIO에 우리가 만든 server를 담아서 이를 통해 실시간 통신을 제어한다.

const moment = require('moment')

io.on('connection',(socket) => {
  socket.on('chatting', (data) => {
    const {name, msg} = data
    io.emit('chatting', {
      name: name,
      msg: msg,
      time: moment(new Date()).format('HH:mm A') // format은 moment에서 제공하는 기능(시간:분)
    })
  })
} )


// app.use() => 디렉토리 내부 파일들을 사용한다.
// path.join => 주소를 합친다(현재 프로젝트 폴더 +  src)
app.use(express.static(path.join(__dirname,'src')))


const PORT = process.env.PORT || 5000
server.listen(PORT, () => {
  console.log(`${PORT} 서버가 동작한다!`)
})

