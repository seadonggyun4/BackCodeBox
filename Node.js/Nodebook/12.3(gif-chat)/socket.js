// [ 웹소켓 서버 ]
// 서버 -> 클라이언트
const SocketIO = require('socket.io');

module.exports = (server) => {
  const io = SocketIO(server, { path: '/socket.io' });

  io.on('connection', (socket) => { // 웹소켓 연결 시
    const req = socket.request;
    const ip = req.headers['x-forwarded-for'] || req.connection.remoteAddress; // 접속자 ip 알아내는 방법
    console.log('새로운 클라이언트 접속!', ip, socket.id, req.ip);
    socket.on('disconnect', () => { // 연결 종료 시
      console.log('클라이언트 접속 해제', ip, socket.id);
      clearInterval(socket.interval);
    });
    socket.on('error', (error) => { // 에러 시
      console.error(error);
    });
    socket.on('reply', (data) => { // 클라이언트로부터 메시지 수신, 직접 만든 이벤트 => 이벤트 명을 동일하게 사용해야 메시지 수신가능
      console.log(data);
    });
    socket.interval = setInterval(() => { // 3초마다 클라이언트로 메시지 전송
      socket.emit('news', 'Hello Socket.IO'); // news 는 이벤트명,  'Hello Socket.IO'는 데이터 => 이를 클라이언트에서 받기 위해서는 socket.on('new',)를 사용해야 한다.
    }, 3000);
  });
};