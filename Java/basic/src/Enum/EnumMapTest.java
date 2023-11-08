package EnumStudy;

import java.util.EnumMap;

// EnumMap은 Enum타입을 키로  사용할수 있도 도와주는 클래스
public class EnumMapTest {
	public static void main(String[] args) {
		EnumMap emap = new EnumMap(Day.class);
		emap.put(Day.SUNDAY, "일요일"); // EnumStudy.Day enum의 SUNDAY를 키값으로 "일요일"이라는 문자열을 할당
		emap.put(Day.THIRSDAY, "목요일");

		System.out.println(emap.get(Day.SUNDAY)); // EnumStudy.Day.SUNDAY를 가져오면 할당된 "일요일" 문자열 출력
	}
}
