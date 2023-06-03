const EventEmitter = require('events')
const emitter = new EventEmitter()



// ellie 이벤트 수신시에 로직
const callBack = (data) => {
    console.log('data :', data)
}
emitter.on('ellie', callBack)


// ellie 이벤트에  데이터 송신
emitter.emit('ellie', { message: 1})
emitter.emit('ellie', { message: 2})
emitter.removeListener('ellie', callBack); // ellie 이벤트에서 콜백함수 제거
emitter.emit('ellie', { message: 3})