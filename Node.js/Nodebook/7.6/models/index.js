const Sequelize = require('sequelize')
const User = require('./user.js')
const Comment = require('./comment.js')

const env = process.env.NODE_ENV || 'development'
const config = require('../config/config.json')[env]
const db = {}

// Sequelize는 시퀼라이즈 패키지 이자 생성자
// config/config.json 에서 데이터베이스 설정을 불러온 후 new Sequelize를 통해 MySol 연결객체 생성
const sequelize = new Sequelize(config.database, config.username, config.password, config)

db.sequelize = sequelize // 연결객체를 재사용하기위해 db객체에 넣어 보관

// db 객체에 User, Comment 모델을 담아둔다 => db 객체를 require하여 두 모델에 접근할 수 있다.
db.User = User
db.Comment = Comment

// static.init 매서드 호출 => init이 실행되어야 테이블이 모델로 연결
User.init(sequelize)
Comment.init(sequelize)

// associate 매서드 => 다른 테이블과의 관계를 연결
User.associate(db)
Comment.associate(db)

module.exports = db