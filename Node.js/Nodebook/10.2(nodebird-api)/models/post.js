// [ 게시글 정보를 저장하는 모델 ]
// 게시글, 이미지 경로를 저장
const Sequelize = require('sequelize');

module.exports = class Post extends Sequelize.Model {
  static init(sequelize) {
    return super.init({
      content: {
        type: Sequelize.STRING(140),
        allowNull: false,
      },
      img: {
        type: Sequelize.STRING(200),
        allowNull: true,
      },
    }, {
      sequelize,
      timestamps: true,
      underscored: false,
      modelName: 'Post',
      tableName: 'posts',
      paranoid: false,
      charset: 'utf8mb4',
      collate: 'utf8mb4_general_ci',
    });
  }

  static associate(db) {
    // post : user => N : 1
    db.Post.belongsTo(db.User);
    // post : hashtag => N : M
    db.Post.belongsToMany(db.Hashtag, { through: 'PostHashtag' });
  }
};