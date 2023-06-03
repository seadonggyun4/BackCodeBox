// buffer
// - 데이터를 "유니코드" 로 변환한다.
// - 버퍼는 배열이다.
// - 배열에 있는 데이터에 접근하면 "아스키 코드" 로 변환이 된다.
const buf = Buffer.from('Hi');

console.log('버퍼 :', buf)
console.log('버퍼 첫번째 :',buf[0])
console.log('버퍼 두번째:',buf[1])


// 버퍼 생성
const buf2 = Buffer.alloc(2) // 초기화된 메모리할당
const buf3 = Buffer.allocUnsafe(2) // 초기화가 안된 메모리 할당(조금더 빠름)
console.log('버퍼2 초기 :',buf2)
console.log('버퍼3 초기 :',buf3)


// Hi문자열 할당후 복제
buf2[0] = 72
buf2[1] = 105
buf2.copy(buf3)
console.log('버퍼2 할당후 :',buf2.toString())
console.log('버퍼3 복제후:',buf3.toString())

// 버퍼 묶기
const newBuf = Buffer.concat([buf, buf2, buf3])
console.log(newBuf.toString())