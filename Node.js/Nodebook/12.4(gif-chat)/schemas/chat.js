const mongoose = require('mongoose')

const { Schema } = mongoose
const { Types: { ObjectId } } = Schema

const chatSchema = new Schema({
    room:{// 채팅방 아이디
        type: ObjectId,
        required: true,
        ref: 'Room'// Room 스키마와 연결
    },
    user:{// 채팅을 한 사람
        type: String,
        required: true,
    },
    chat:{// 채팅 내역
        type: String
    },
    gif:{// 이미지 주소
        type: String
    },
    createdAt:{// 채팅 시간
        type: Date,
        default: Date.now
    }
})

module.exports = mongoose.model('Chat', chatSchema)