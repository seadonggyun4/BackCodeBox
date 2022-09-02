// MYsql 의 테이블은 시퀼라이즈의 모델과 대응된다. => 시퀼라이즈는 모델과 테이블을 연결지어주는 기능을 한다.
// 시퀼라이즈는 기본적으로 모델 이름은 단수형으로, 테이블 이름은 복수형으로 사용합니다.

const Sequelize = require('sequelize')


// User 모델 생성 => 시퀼라이즈 모델을 확장한 클래스로 선언
// 모델은 크게 static init, static associate로 분류된다.
// static init => 테이블에 대한 설정
// static associate => 다른 모델과의 관계 설정
module.exports = class User extends Sequelize.Model {
    static init(sequelize) {
        return super.init({
            name: {
                type: Sequelize.STRING(20),
                allowNull: false,
                unique: true
            },
            age: {
                type: Sequelize.INTEGER.UNSIGNED,
                allowNull: false
            },
            married: {
                type: Sequelize.BOOLEAN,
                allowNull: false
            },
            comment: {
                type: Sequelize.TEXT,
                allowNull: true
            },
            created_at : {
                type: Sequelize.DATE,
                allowNull: false,
                defaultValue: Sequelize.NOW
            },
        },{
            sequelize, 
            timestamps: false,
            underscored: false,
            modelName: 'User',
            tableName: 'Users',
            paranoid: false,
            charset: 'utf8',
            collate: 'utf8_general_ci'
        });
    }

    static associate (db){
        db.User.hasMany(db.Comment, {foreignKey: 'commenter', sourceKey: 'id'}) // 1:N 관계(user:comments)정의에서 1 인 user테이블은 .hasMany()와 같은 메서드를 사용해 Comment와 연결짓는다. 
    }
}

