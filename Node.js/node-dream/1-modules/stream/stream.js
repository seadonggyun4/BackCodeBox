const fs = require('fs')


const data = []

// [ 읽고싶은 파일의 스트림 생성 ]
// - 옵셔널 체이닝으로 각 이벤트에 따른 로직을 따로 설정 가능
const readStream = fs.createReadStream('./file.txt', {
    highWaterMark: 8, // base: 64 kbytes
    encoding: 'utf-8',
})


readStream.on('data', (chunk) => {
    data.push(chunk)
    console.count('데이터 받아오는 중!')
})

readStream.on('end', () => {
    console.log(data.join(''))
})

readStream.on('error', (e) => {
    console.error(e)
})