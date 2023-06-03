const fs = require('fs')

// fs를 사용할수 있는 3가지 방법!!!


// [ 동기적 방법: renameSync ]
// - 에러 발생시 node서버 자체가 죽어버린다. 이를 방지하기 위해서는 try{}catch(e){} 문으로 에러상황일때의 대응을 해주어야 한다.
// try{
//     fs.renameSync('./text.txt', './text-new.txt')
// } catch (e){
//     console.error(e)
// }


// [ 비동기적 방법: rename ]
// - 콜백함수 내부의 로직은 에러상황, 정상상황 둘다 실행된다.
fs.rename('./text-new.txt', './text.txt', (e) => {
    console.log(e)
})


// [ 비동기적 방법: promise rename ]
// - then, catch 로 콜백로직을 구분지어 실행할 수 있다.
fs.promises
    .rename('./text.txt', './text-new.txt')
    .then(() => {
        console.log('정상 동작')
    })
    .catch(() => {
        console.log('에러발생')
    })
