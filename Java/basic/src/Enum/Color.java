package EnumStudy;

// enum 은 interface를 받을수도 있다.
public enum Color implements Printer {
	RED("FF0000"),
	GREEN("00FF00"),
	BLUE("0000FF");

	private String rgb;

	private Color(String rgb) {
		this.rgb = rgb;
	}

	@Override
	public void print() {
		System.out.println("rgb: " + rgb);
	}
}
