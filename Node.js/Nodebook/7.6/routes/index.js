// [GET / 요청했을 때의 라우터]
const express = require('express');
const User = require('../models/user');

const router = express.Router();

router.get('/', async (req, res, next) => {
  try {
    const users = await User.findAll(); // User.findAll() 매서드로 모든 user모듈의 데이터를 찾아온다 
    res.render('sequelize', { users }); // sequelize.html을 렌더링할 때 결과값인 users을 넣습니다.
  } catch (err) {
    console.error(err);
    next(err);
  }
});

module.exports = router;