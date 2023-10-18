import java.util.HashMap;
import java.util.Map;

public class MapExam {
    public static void main(String[] args) {
        // 맵은 키, 값으로 이루어진 클래스 이다.
        Map<String, String> map = new HashMap<>();
        map.put("k1", "hello");
        map.put("k2", "hi");
        map.put("k3", "test");
        map.put("k3", "안녕하세요!!!!"); // 기존에 같은키로 선언된 인스턴스를 덮어쓴다.

        System.out.println(map.get("k1"));
        System.out.println(map.get("k2"));
        System.out.println(map.get("k3"));
    }
}
