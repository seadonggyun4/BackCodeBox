// 현재 동작하고 있는 프로그램의 환경 변수
const process = require('process')


console.log(process.env)

// 테스크 큐에 다른 콜백함수들이 있더라도, 무시하고 가장먼저 수행된다.
process.nextTick(() => {
    console.log('nextTick')
})