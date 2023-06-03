// os에 대한 정보를 담고 있는 객체
const os = require('os')

// 운영체제별 줄바꿈
// window는 \r\n, 맥은 \n
console.log(os.EOL === '\n')

console.log(os.cpus())