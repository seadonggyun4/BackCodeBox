import java.util.HashMap;
import java.util.Iterator;
import java.util.Map;
import java.util.Set;

public class MapExam02 {
    public static void main(String[] args) {
        // 맵은 키, 값으로 이루어진 클래스 이다.
        Map<String, String> map = new HashMap<>();
        map.put("k1", "hello");
        map.put("k2", "hi");
        map.put("k3", "test");


        Set<String> keySet = map.keySet(); //맵 변수내부의 키값을 저장
        Iterator<String> iterator = keySet.iterator();
        while (iterator.hasNext()){
            String key = iterator.next();
            String value = map.get(key);
            System.out.println(value);
        }
    }
}
