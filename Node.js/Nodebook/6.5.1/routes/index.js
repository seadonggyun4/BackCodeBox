const express = require('express');

const router = express.Router();

// GET / 라우터
router.get('/', (req, res) => {
  //render는 express가 템플릿 렌더링을 하기위해 추가한 매서드  
  res.render('index', { title: 'Express' }); // => index.pug를 html로 랜더링하면서 {title: "Express"} 라는 객체를 변수로 집어 넣는다.
});

module.exports = router;

router.get('/abc', (req, res) => {
  res.send('GET /abc');
});
router.post('/abc', (req, res) => {
  res.send('POST /abc');
});

router.route('/abc')
  .get((req, res) => {
    res.send('GET /abc');
  })
  .post((req, res) => {
    res.send('POST /abc');
  });