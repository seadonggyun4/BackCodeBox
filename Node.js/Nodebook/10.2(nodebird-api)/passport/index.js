const passport = require('passport');
const local = require('./localStrategy');
const kakao = require('./kakaoStrategy');
const User = require('../models/user');

module.exports = () => {
  // serializeUser => 로그인시 실행되며 req.session 객체에 어떤 데이터를 저장할지 정하는 메서드
  passport.serializeUser((user, done) => {
    done(null, user.id); // 세션에 사용자 아이디만 저장
  });

  // deserializeUser => 매 요청시 실행되는 매서드
  // serializeUser 내부 done()매서드의 두번째 인자를 매개변수로 받는다.
  // 매개변수로 받은 값(id)를 통해 데이터베이스에서 사용자 정보를 조회한다.
  // req.user에는 팔로워와 팔로잉 목록도 같이 불려오게 된다.
  passport.deserializeUser((id, done) => {
    User.findOne({
      where: { id },
      include: [{
        model: User,
        attributes: ['id', 'nick'],
        as: 'Followers',
      }, {
        model: User,
        attributes: ['id', 'nick'],
        as: 'Followings',
      }],
    })
      .then(user => done(null, user))
      .catch(err => done(err));
  });

  local();
  kakao();
};

/*
 [ serializeUser,deserializeUser 매서드 들 사용 이유 ]
 - serializeUser는 사용자 정보 객체를 세션에 아이디로 저장, deserializeUser는 세션에 저장한 아이디를 통해 사용자 정보 객체를 불러온다.
 - 세션에 불필요한 데이터를 담아두지 않기 위한 과정
*/