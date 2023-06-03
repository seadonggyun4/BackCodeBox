const path = require('path')

// 실행되는 폴더의 경로
// /Users/dgseo/Desktop/BackCodeBox/Node.js/node-dream/1-modules
console.log(__dirname)

// 실행되는 파일의 경로
// /Users/dgseo/Desktop/BackCodeBox/Node.js/node-dream/1-modules/path.js
console.log(__filename)


// 파일 이름
console.log(path.basename(__filename))


// join 을 활용한 경로 설정
console.log(path.join(__dirname, 'image'))


// 경로와 같은것들은 운영체제별로 많이 달라질수 있으니, 항상 하드코딩이 아닌 계산된 값을 바탕으로 작성!!!!