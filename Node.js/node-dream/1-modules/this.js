// 함수에서의 this는 global이다
function functionThis(){
    console.log('function :', this === global)
}
functionThis()


// 클래스에서의 this는 클래스 내부에 선언된 변수이다.
class A{
    constructor(num) {
        this.num = num
    }
    memberFunction(){
        console.log('class :', this)
    }
}
const a = new A(1)
a.memberFunction()