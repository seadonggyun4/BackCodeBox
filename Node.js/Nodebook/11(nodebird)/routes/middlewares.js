// [ 라우터에 접근 권한을 제어하는 미들웨어 ]
// passport는 req객체에 isAuthenticated 매서드를 추가 => 로그인 중이면 true, 아니면 false

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
      const message = encodeURIComponent('로그인한 상태입니다.');
      res.redirect(`/?error=${message}`);
    }
  };