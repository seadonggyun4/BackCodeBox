const express = require('express');
const jwt = require('jsonwebtoken');

const { verifyToken } = require('./middlewares');
const { Domain, User } = require('../models');

const router = express.Router();

// [ 토큰을 발급하는 라우터 ]
// 전달받은 클라이언트 비밀 키로 도메인이 등록된 것인지를 확인 => 등록되지 않은 도메인 이라면 에러메시지, 등록된 도메인이면 토큰 발급
router.post('/token', async (req, res) => {
  const { clientSecret } = req.body;
  try {
    const domain = await Domain.findOne({
      where: { clientSecret },
      include: {
        model: User,
        attribute: ['nick', 'id'],
      },
    });
    if (!domain) {
      return res.status(401).json({
        code: 401,
        message: '등록되지 않은 도메인입니다. 먼저 도메인을 등록하세요',
      });
    }
    //토큰 발급 => jwt.sign(토큰내용, 토큰 비밀키, 토큰 설정)
    const token = jwt.sign({
      id: domain.User.id,
      nick: domain.User.nick,
    }, process.env.JWT_SECRET, {
      expiresIn: '1m', // 1분
      issuer: 'nodebird',
    });
    return res.json({
      code: 200,
      message: '토큰이 발급되었습니다',
      token,
    });
  } catch (error) {
    console.error(error);
    return res.status(500).json({
      code: 500,
      message: '서버 에러',
    });
  }
});

// [ 토큰을 테스트해볼 수 있는 라우터 ]
router.get('/test', verifyToken, (req, res) => {
  res.json(req.decoded);
});

module.exports = router;