const logger = require('./logger.js')
const Logger = new logger.Logger('드림코더')

// Logger 클래스 생성 및 EventEmitter로 상속하고 있는 이벤트 "log" 수신
Logger.on('log', (e) => {
    console.log(e)
})


// Logger 클래스에 포함된 함수 사용
Logger.makeHi(() => {
    console.log('콜백함수 : 만나서 반가워!!')
})

Logger.makeGoodBy(() => {
    console.log('콜백함수 : 잘가~~~')
})


