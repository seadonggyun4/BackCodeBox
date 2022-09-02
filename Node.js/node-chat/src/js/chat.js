// [ 프론트 엔드]
"use strict"

// io() 객체를 받는다 -> 이벤트 핸들링을 위한 socket 객체
const socket = io();

const nickname = document.querySelector('#nickname')
const chatList = document.querySelector('.chatting-list')
const chatInput = document.querySelector('.chatting-input')
const sendButton = document.querySelector('.send-button')
const displayContainer = document.querySelector('.display-container')
const chattingList = document.querySelector('.chatting-list')

function send(){
  const param = {
    name: nickname.value,
    msg: chatInput.value,
  }

  socket.emit('chatting', param)
}


chatInput.addEventListener('keypress', (event) => {
  if(event.keyCode === 13){
    send()
  }
})


sendButton.addEventListener('click', send)




socket.on('chatting', (data) => {
  const {name, msg, time} = data
  const item = new LiModel(name, msg, time)
  item.makeLi()
  chattingList.scrollTo(0, chattingList.scrollHeight)
})

function LiModel(name, msg, time){
  this.name = name
  this.msg = msg
  this.time = time

  this.makeLi = () => {
    const li = document.createElement('li')
    li.classList.add(nickname.value === this.name ? 'sent' : 'received')
    const dom = `<span class="profile">
      <span class="user">${this.name}</span>
      <img class="image" src="https://Placeimg.com/50/50/any" alt="user-image">
      </span>
      <span class="message">${this.msg}</span>
      <span class="time">${this.time}</span>`
    li.innerHTML = dom
    chatList.appendChild(li)
  }
}