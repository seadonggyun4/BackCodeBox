// [ 댓글에 관련된 CRUD를 진행하는 라우터: POST /comments, PATCH /comments/:id, DELETE /comments/:id ]
const express = require('express');
const Comment = require('../schemas/comment');

const router = express.Router();

// 다큐먼트를 등록하는 라우터: 댓글저장
router.post('/', async (req, res, next) => {
  try {
    const comment = await Comment.create({
      commenter: req.body.id,
      comment: req.body.comment,
    });
    console.log(comment);
    const result = await Comment.populate(comment, { path: 'commenter' });
    res.status(201).json(result);
  } catch (err) {
    console.error(err);
    next(err);
  }
});


router.route('/:id')
  .patch(async (req, res, next) => { // 다큐먼트를 수정하는 라우터: 댓글수정
    try {
      const result = await Comment.update({
        _id: req.params.id,
      }, {
        comment: req.body.comment,
      });
      res.json(result);
    } catch (err) {
      console.error(err);
      next(err);
    }
  })
  .delete(async (req, res, next) => { // 다큐먼트를 삭제하는 라우터: 댓글 삭제
    try {
      const result = await Comment.remove({ _id: req.params.id });
      res.json(result);
    } catch (err) {
      console.error(err);
      next(err);
    }
  });

module.exports = router;