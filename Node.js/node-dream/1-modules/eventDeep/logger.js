const EventEmitter = require('events');


// Logger 클래스 생성 및 EventEmitter 상속
// - 클래스 상속은 기존 클래스(부모 클래스)의 속성과 메서드를 다른 클래스(자식 클래스)에게 상속하고, 자식 클래스는 상속된 멤버들을 사용할 수 있게 합니다. 자식 클래스는 부모 클래스의 특성을 확장하거나 변경하여 사용할 수 있습니다.
class Logger extends EventEmitter {
    constructor(name) {
        super(); // 부모클래스 생성자 호출, extends일때는 항시 사용
        this.name = name
    }

    makeHi(callback){
        this.emit('log', `Hi😄 ${this.name}`);
        callback();
    }

    makeGoodBy(callback){
        this.emit('log', `GoodBy🥹 ${this.name}`);
        callback();
    }
}


module.exports.Logger = Logger;