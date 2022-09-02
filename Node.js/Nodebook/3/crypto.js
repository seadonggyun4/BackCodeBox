const crypto = require('crypto')

// crypto.createHash('sha512') : 512알고리즘을 통해 암호화
// .update() : 암호화할 문자열
// .digest() : 인코딩할 알고리즘
console.log('base64 :', crypto.createHash('sha512').update('비밀번호').digest('base64'))

// randomBytes() => **바이트 길이의 문자열을 만듭니다, 아래는 64바이트 문자열
// buf가 랜덤함수를 통해 생긴 문자열의 값
// .toString('base64') => base64알고리즘을 통해 암호화
crypto.randomBytes(64, (err, buf) => {
    const salt = buf.toString('base64')
    console.log('salt :', salt)
    //.pbkdf2() : 비민번호, salt, 반복횟수, 출력바이트, 해시알고리즘, (결과값) 순서이다.
    crypto.pbkdf2('비밀번호', salt, 100000, 64, 'sha512', (err,key) => {
        console.log('password:', key.toString('base64'))
    })
})