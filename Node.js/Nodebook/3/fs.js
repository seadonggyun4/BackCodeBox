// [ 파일읽는 작업은 비동기로 처리되게 때문에 여러번의 I/O작업이 발생할때 순서대로 되진 않는다. ]
const fs = require('fs').promises

// console.log('시작')

// fs.readFile('./readme2.txt')
// .then((data) => {
//     console.log('1번: ', data.toString())
// })
// .catch((err) => {
//     console.log(err)
// })

// fs.readFile('./readme2.txt')
// .then((data) => {
//     console.log('2번: ', data.toString())
// })
// .catch((err) => {
//     console.log(err)
// })

// fs.readFile('./readme2.txt')
// .then((data) => {
//     console.log('3번: ', data.toString())
// })
// .catch((err) => {
//     console.log(err)
// })

// console.log('끝')




// [ 순서대로 짜는 방법 ] 
console.log('시작')

fs.readFile('./readme2.txt')
.then((data) => {
    console.log('1번 :', data.toString())
    return fs.readFile('./readme2.txt')
})
.then((data) => {
    console.log('2번', data.toString())
    return fs.readFile('./readme2.txt')
})
.then((data) => {
    console.log('3번 :', data.toString())
    console.log('끝')
})
.catch((err) => {
    console.log(err)
})


