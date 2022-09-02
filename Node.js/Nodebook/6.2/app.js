/////////////////////////////////////////////////////////////////////////

const express = require('express');
const morgan = require('morgan');
const cookieParser = require('cookie-parser');
const session = require('express-session');
const dotenv = require('dotenv');
const path = require('path');

dotenv.config();// .env 파일을 읽어 process.env로 만듭니다.
const app = express();
app.set('port', process.env.PORT || 3000);

app.use(morgan('dev'));// 요청과 응답에 대한 정보를 콘솔에 기록합니다.
app.use('/', express.static(path.join(__dirname, 'public')));
app.use(express.json());// body-parser사용
app.use(express.urlencoded({ extended: false }));// body-parser사용: urlencoded는 주소형식으로 데이터를 보내는 방식
app.use(cookieParser(process.env.COOKIE_SECRET));// 쿠키를 해석해 req.cookies 객체로 만든다. 
app.use(session({// 세션관리용 미들웨어로 특정 사용자를 위한 데이터를 임시적으로 저장해둘 때 매우 유용합니다.
  resave: false,
  saveUninitialized: false,
  secret: process.env.COOKIE_SECRET,
  cookie: {
    httpOnly: true,
    secure: false,
  },
  name: 'session-cookie',
}));

const multer = require('multer');
const fs = require('fs');

try {
  fs.readdirSync('uploads');
} catch (error) {
  console.error('uploads 폴더가 없어 uploads 폴더를 생성합니다.');
  fs.mkdirSync('uploads');
}
const upload = multer({
  storage: multer.diskStorage({
    destination(req, file, done) {
      done(null, 'uploads/');
    },
    filename(req, file, done) {
      const ext = path.extname(file.originalname);
      done(null, path.basename(file.originalname, ext) + Date.now() + ext);
    },
  }),
  limits: { fileSize: 5 * 1024 * 1024 },
});
app.get('/upload', (req, res) => {
  res.sendFile(path.join(__dirname, 'multipart.html'));
});
app.post('/upload', upload.single('image'), (req, res) => {
  console.log(req.file);
  res.send('ok');
});

app.get('/', (req, res, next) => {
  console.log('GET / 요청에서만 실행됩니다.');
  next();// 다음 미들웨어를 실행하는 함수, 이를 사용하지 않으면 다음 미들웨어로 넘어가지 않는다.
}, (req, res) => {
  throw new Error('에러는 에러 처리 미들웨어로 갑니다.')
});
app.use((err, req, res, next) => {
  console.error(err);
  res.status(500).send(err.message);
});

app.listen(app.get('port'), () => {
  console.log(app.get('port'), '번 포트에서 대기 중');
});


/*
    [익스프레스 미들웨어]
    - 미들웨어는 익스프레스 의 핵심입니다.
    - 요청과 응답 중간에 위치하여 로직을 처리합니다.
    - 미들웨어는 요청과 응답을 조작하여 기능을 추가하기도 하고, 나쁜요청을 걸러내기도 앟ㅂ니다
    
    - app.use(미들웨어)
    - app.use((req, res, next) => {})
    - app.use(주소, 함수) : 주소를 첫번째 인수로 넣어주면 미들웨어는 해당 주소에서만 실행된다.

    - app.use(), app.get()같은 라우터에 미들웨어를 여러 개 장착할 수 있습니다.
*/