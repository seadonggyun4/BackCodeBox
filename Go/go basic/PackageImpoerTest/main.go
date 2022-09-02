package main

import "kerygma.com/greeting"

// import될 패키지는 모듈파일에서 연결해야 한다.
// main mod를 만든후 mod 내부에 명령어를 추가해야한다
/*

mod 에  replace kerygma.com/greeting => ../greeting
명령어를 통해도메인 주소를 컴퓨터 내부의 디렉토리로 연결시킨다.

그 후 터미널에 go get kerygma.com/greeting
*/
func main() {
	greeting.SayHello()
}
