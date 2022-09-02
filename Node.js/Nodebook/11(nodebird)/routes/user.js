// [ 다른사람을 팔로우 하는 기능의 라우터 ]
const express = require('express');

const { isLoggedIn } = require('./middlewares');
const User = require('../models/user');

const router = express.Router();

router.post('/:id/follow', isLoggedIn, async (req, res, next) => {
  try {
      // 팔로우할 사용자를 데이터베이스에서 조회
    const user = await User.findOne({ where: { id: req.user.id } });
    if (user) {
      // 시퀄라이즈에 추가한 addFollowing 매서드로 현재 로그인한 사용자 와의 관계 지정
      await user.addFollowing(parseInt(req.params.id, 10));
      res.send('success');
    } else {
      res.status(404).send('no user');
    }
  } catch (error) {
    console.error(error);
    next(error);
  }
});

module.exports = router;