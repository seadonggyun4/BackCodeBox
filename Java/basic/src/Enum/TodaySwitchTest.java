package EnumStudy;

public class TodaySwitchTest {
	public static void main(String[] args) {
		Day day = Day.SUNDAY;

		// switch 문을 작성할때 Day가 가지고있는 상수명으로 케이스가 작성이 되어야 한다. -> 그 이외의 케이스는 컴파일에러 발생
		switch (day) {
			case SUNDAY:
				System.out.println("Sunday");
				break;
			case THUSDAY:
				System.out.println("Thusday");
				break;
			default:
				System.out.println("그 외의 요일");
		}
	}
}
