public class Person {
	/*
	 * 스테틱이 붙은 필드, 클래스는 JVM에서 바로 클래스로 참조하고 인스턴스를 만들지 않고 사용 가능하도록 메모리 영역에 올려둔다
	 * - 인스턴스를 만들지 않고 클래스 이름으로 바로 참조해서 사용 가능하다.
	 * - 장점: 인스턴스를 할당하지 않아 메모리 절략이 가능하다.
	 * - 주의할점: 클래스 필드, 함수를 인스턴스를 할당하더라도 사용하려는 필드, 함수는 각각 할당된 인스턴스 별로 생성되는게 아니다.
	 * */

	String name; // 인스턴스 필드
	String address;
	boolean isVip;

	static int count = 0; // 클래스 필드

	static {

		count = 100; // 클래스 필드는 static 블록 단위에서 초기화 할 수 있다.
	}

	public void printName() { // 인스턴스 메서드
		System.out.println("Name :" + name);

	}

	public static void printCount() { // 클래스 메서드
		// static 한 클래스 메서드에서는 인스턴스 필드,메서드를 사용할수 없다.
		System.out.println("Count" + count);
	}
}
