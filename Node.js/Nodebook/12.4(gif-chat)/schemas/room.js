const mongoose = require('mongoose')

const { Schema } = mongoose
const roomSchema = new Schema({
    title:{ //방 제목
        type: String,
        required: true,
    },
    max: {// 최대 수용 인원
        type:Number,
        required:true,
        default: 10,
        min: 2,
    },
    owner: {// 방장
        type:String,
        required:true
    },
    password: {// 비밀번호
        type:String
    },
    createdAt: {// 생성 시간
        type:Date,
        default:Date.now
    }
})

module.exports = mongoose.model('Room', roomSchema)