const fs = require('fs')


const writeStream = fs.createWriteStream('./file3.txt')
writeStream.on('finish', () => {
    console.log('파일생성 끝!')
})


writeStream.write('hello!')
writeStream.write('\nworld ~~~~')
writeStream.end()