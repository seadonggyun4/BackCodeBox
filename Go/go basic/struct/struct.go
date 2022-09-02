package main // 구조체 내부의 구조체
import "fmt"

//대문자 선언 시 모두 노출 가능

type User struct {
	Name string // 현재 패키지를 벗어나 외부에 노출이 된다.
	age  int    // 외부에 알려지지 않음
	Id   string
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	price    int
}

func main() {
	normalUser := User{Name: "이순신", age: 55, Id: "Lee"}
	vipUser := VIPUser{User{Name: "이순신", age: 55, Id: "Lee"}, 5, 100000}
	fmt.Println(normalUser, vipUser)
	fmt.Printf("유저이름 : %s, 나이 : %d, 아이디 : %s \n", normalUser.Name, normalUser.age, normalUser.Id)
	fmt.Printf("VIP유저이름 : %s, ", vipUser.UserInfo.Name)

	vipUser2 := vipUser //구조체 복사가 가능하다.
	fmt.Println(vipUser2)
}
