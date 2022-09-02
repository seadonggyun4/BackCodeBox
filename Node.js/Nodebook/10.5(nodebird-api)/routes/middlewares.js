// [ 라우터에 접근 권한을 제어하는 미들웨어 ]
// passport는 req객체에 isAuthenticated 매서드를 추가 => 로그인 중이면 true, 아니면 false


const jwt = require('jsonwebtoken');

exports.isLoggedIn = (req, res, next) => {
  if (req.isAuthenticated()) {
    next();
  } else {
    res.status(403).send('로그인 필요');
  }
};

exports.isNotLoggedIn = (req, res, next) => {
  if (!req.isAuthenticated()) {
    next();
  } else {
    res.redirect('/');
  }
};

exports.verifyToken = (req, res, next) => {
  try {
    // 요청 헤더에 저장된 토큰을 사용합니다.
    // jwt.verify()메서드로 토큰을 검증할 수 있습니다. 
    req.decoded = jwt.verify(req.headers.authorization, process.env.JWT_SECRET);
    // 인증에 성공하면 req.decoded을 발급받는다 => v1.js 에서 내부 내용물 사용가능
    return next();
  } catch (error) {
    if (error.name === 'TokenExpiredError') { // 유효기간 초과
      return res.status(419).json({
        code: 419,
        message: '토큰이 만료되었습니다',
      });
    }
    return res.status(401).json({
      code: 401,
      message: '유효하지 않은 토큰입니다',
    });
  }
};