// [mysql 데이터베이스 시퀼라이즈 와 연결]
const Sequelize = require('sequelize');
const env = process.env.NODE_ENV || 'development';
const config = require('../config/config')[env];
const User = require('./user');
const Post = require('./post');
const Hashtag = require('./hashtag');

const db = {}; // 데이터를 담을 데이터 객체 생성
const sequelize = new Sequelize(
  config.database, config.username, config.password, config,
); 

// 각각의 모델들을 시퀼라이즈 객체에 연결
db.sequelize = sequelize;
db.User = User;
db.Post = Post;
db.Hashtag = Hashtag;

User.init(sequelize);
Post.init(sequelize);
Hashtag.init(sequelize);

User.associate(db);
Post.associate(db);
Hashtag.associate(db);

module.exports = db;