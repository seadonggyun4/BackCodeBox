package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//시도 횟수 변수 ,점수 변수를 생성
	var try int
	var score int

	try = 0
	score = 0

	print_to_Dest()
	// 무한 반복문으로 돌린다.
	for {
		//랜덤 숫자 생성 하는 함수 호출
		no1, no2 := randNumbers()
		// 함수를 변수로 받아 랜덤으로 식을 내는 코드를 짜야됨.
		fmt.Printf("%d X %d의 값은: \n", no1, no2)
		//유저가 답을 입력하는 함수를 호출
		res := no1 * no2
		answer := inputUserAnswer()

		if check_to_Answer(answer, res) {
			fmt.Println("정답입니다.")
			fmt.Println("1점을 얻었습니다.")
			score++
			try++
			fmt.Printf("현재점수 : %d \n", score)
		} else {
			fmt.Println("오답입니다.")
			fmt.Println("1점이 감소됩니다.")
			score--
			try++
			fmt.Printf("현재점수 : %d \n", score)
		}
		if score == 5 {
			fmt.Println("축하합니다! 점수가 5점을 넘으셨습니다!!!!")
			break
		}

	}
	fmt.Printf("총 도전횟수는 ? : %d 번", try)
}

func randNumbers() (int, int) {
	//시드는 언제나 바뀌게
	rand.Seed(time.Now().UnixNano()) //
	x := rand.Intn(10)
	y := rand.Intn(10)
	//x,y 모두 1~10까지 숫자중 하나로 랜덤 선출

	return x, y
}

func inputUserAnswer() int {
	fmt.Println("정답은 ? : ")
	reader := bufio.NewReader(os.stdin)
	line, _ := reader.ReadString('\n')
	line = string.TrimSpace(line)
	answer, _ := strconv.Atoi(line)

	return answer
}

func check_to_Answer(x, y int) bool {
	if x == y {
		return true
	} else {
		return false
	}
}

func print_to_Dest() {
	fmt.Println("구구단 게임을 시작합니다.")
	fmt.Println("랜덤 숫자를 2개를 컴퓨터가 정한 뒤 유저가 답을 입력하는 방식입니다.")
	fmt.Println("정답을 맞추면 1점을 얻고, 정답이 틀리면 1점이 감소가 됩니다.")
	fmt.Println("5점을 얻으면, 게임이 종료!!!")
}
