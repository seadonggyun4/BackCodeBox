package Enum;

// enum 도 생성자를 가질수 있다.
// enum 도 메소드나 변수를 가질수 있다.
public enum Gender {
	MALE("XY"),
	FEMALE("XX");

	private String chromosome;

	private Gender(String chromosome) {
		this.chromosome = chromosome;
	}

	// enum의 타입(MALE, FEMALE)이 아닌 생성자로 선언된 인스턴스의 값을 뽑아오고 싶다면 toString 매서드를 오버라이딩 한다.
	@Override
	public String toString() {
		return super.toString() + ": " + chromosome;
	}

	public void print() {
		System.out.println("염색체 " + super.toString() + "정보 :" + chromosome);
	}
}
