// [ 몽고DB user 컬렉션에 대응되는 userSchema를 만든다. ]

const mongoose = require('mongoose');

const { Schema } = mongoose; // mongoose 모듈의 Schema 생성자를 사용해 스키마를 만든다.
const userSchema = new Schema({
  name: {
    type: String,
    required: true,
    unique: true,
  },
  age: {
    type: Number,
    required: true,
  },
  married: {
    type: Boolean,
    required: true,
  },
  comment: String,
  createdAt: {
    type: Date,
    default: Date.now,
  },
});

module.exports = mongoose.model('User', userSchema);