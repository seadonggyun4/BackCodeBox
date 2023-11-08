package Enum;

// 클래스 Today의 인스턴스 today를 선언
// today로 Today의 getter, setter 를 접근할때 enum 으로 할당가능한 범위가 지정된 EnumStudy.Day 에서만 값을 set 할수 있다.
public class TodayTest {
	public static void main(String[] args) {
		Today today = new Today();
		today.setDay(Day.SUNDAY);
		System.out.println(today.getDay());
	}
}
