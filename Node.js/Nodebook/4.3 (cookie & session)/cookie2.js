const http = require('http');
const fs = require('fs').promises;
const url = require('url');
const qs = require('querystring');

// parseCookies함수 => 문자열을 객체로 바꿔준다.
const parseCookies = (cookie = '') =>
  cookie
    .split(';')
    .map(v => v.split('='))
    .reduce((acc, [k, v]) => {
      acc[k.trim()] = decodeURIComponent(v);
      return acc;
    }, {});

http.createServer(async (req, res) => {
  const cookies = parseCookies(req.headers.cookie); // { mycookie: 'test' }
  // 주소가 /login으로 시작하는 경우
  if (req.url.startsWith('/login')) {
    const { query } = url.parse(req.url);
    const { name } = qs.parse(query);
    const expires = new Date();
    // 쿠키 유효 시간을 현재시간 + 5분으로 설정
    expires.setMinutes(expires.getMinutes() + 5);
    res.writeHead(302, {
      Location: '/',
      'Set-Cookie': `name=${encodeURIComponent(name)}; Expires=${expires.toGMTString()}; HttpOnly; Path=/`,
    });
    res.end();
  // name이라는 쿠키가 있는 경우
  } else if (cookies.name) {
    res.writeHead(200, { 'Content-Type': 'text/plain; charset=utf-8' });
    res.end(`${cookies.name}님 안녕하세요`);
  } else { //name쿠키가 없는 경우 cookie2.html 전송
    try {
      const data = await fs.readFile('./cookie2.html');
      res.writeHead(200, { 'Content-Type': 'text/html; charset=utf-8' });
      res.end(data);
    } catch (err) {
      res.writeHead(500, { 'Content-Type': 'text/plain; charset=utf-8' });
      res.end(err.message);
    }
  }
})
  .listen(8084, () => {
    console.log('8084번 포트에서 서버 대기 중입니다!');
  });


  /*
    [ writeHead() 의 'Set-Cookie' 일 경우에 사용할 수 있는 키,값 쌍 ]

    쿠키명=쿠키값 : 기본적인 쿠키 값
    Expires=날짜 : 만료 기한입니다. 이 기간이 지나면 쿠키가 만료합니다.
    Max-age=초 : Expires와 비슷하지만 날짜 대신 초를 입력할 수 있습니다. 해당 초가 지나면 쿠키가 제거 됩니다.
    Domain=도메인명 : 쿠키가 전송될 도메인을 특정할 수 있습니다.
    Path=URL : 쿠키가 전송될 URL을 특정할 수 있습니다. 기본값은 '/'이고, 이 경우 모든 URL에서 쿠키가 전송 가능
    Secure : HTTPS일 경우에만 쿠키가 전송됩니다.
    HttpOnly: 설정 시 자바스크립트에서 쿠키에 접근할 수 없습니다. 쿠키 조작을 방지하기 위해 설정
  */