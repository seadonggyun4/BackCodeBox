const fs = require('fs')
const zlib = require('zlib')

const readStream = fs.createReadStream('./file.txt'); // read
const zlibStream = zlib.createGzip(); // 압축
const writeStream = fs.createWriteStream('./file4.zip'); // write

writeStream.write('hello!')
writeStream.write('\n데이터 파이핑 테스트!! \n')


// [ 데이터 파이핑 ]
// - readStream 에 writestream 을 연결
const piping = readStream.pipe(zlibStream).pipe(writeStream)


piping.on('finish', () => {
    console.log('done!!!')
})