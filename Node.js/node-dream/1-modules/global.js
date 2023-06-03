// global은 브라우저의 window 와 같은 내장 객체
global.hello = () => {
    console.log('글로벌 객체 공부중');
}


// hello로 정의된 함수 사용
global.hello();
hello();



// [ Console ]
console.log('출력');
console.info('출력');
console.warn('출력');
console.error('출력');
console.table({name: '테스트', state: '안전'});
console.dir({name: '테스트', state: { error: 'red', warn:'yellow' }}, {showHidden: true, colors: false, depth: 2})
console.assert(true, '참 일때 출력')
console.time('for loop')
for(let i=0; i<10; i++){
    i++
}
console.timeEnd('for loop')
// console.clear();