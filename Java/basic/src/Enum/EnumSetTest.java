package Enum;

import java.util.EnumSet;
import java.util.Iterator;

public class EnumSetTest {
	public static void main(String[] args) {
		EnumSet eset = EnumSet.allOf(Day.class); // EnumSet을 통해 EnumStudy.Day enum이 가진 모든 상수를 eset에 넣어준다.

		Iterator<Day> dayIter = eset.iterator();

		// iterator 를 활용해 EnumStudy.Day enum 의 타입을 반복 출력
		while (dayIter.hasNext()) {
			Day day = dayIter.next();
			System.out.println(day);
		}

		System.out.println("=====================================================================");

		EnumSet eset2 = EnumSet.range(Day.SUNDAY,
			Day.THUSDAY); // EnumSet을 통해 EnumStudy.Day enum이 가진 상수중 EnumStudy.Day.SUNDAY 부터 EnumStudy.Day.THUSDAY까지의 값을 eset2에 넣어준다.
		Iterator<Day> dayIter2 = eset2.iterator();
		while (dayIter2.hasNext()) {
			Day day = dayIter2.next();
			System.out.println(day);
		}
	}
}
