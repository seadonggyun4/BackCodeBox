import java.util.ArrayList;
import java.util.Collection;
import java.util.Iterator;

public class ListExam03 {
    public static void main(String[] args){
        // ArrayList 는 Collection의 자식 인터페이스 -> 참조타입이 가진 메서드만 사용가능
        // collection은 내부 값의 순서를 기억 x
        // ArrayList는 내부 값의 순서를 기억 o
        // 타입은 인터페이스로 선언하고 주체는 인스턴스로(클래스) 선언
        Collection<String> collection = new ArrayList<>();
        collection.add("kim");
        collection.add("lee");
        collection.add("hong");



        // Collection이 가지고 있는 매서드만 호출 가능
        System.out.println(collection.size());

        Iterator<String> iter = collection.iterator(); // iterator 는 collection 이 가진 모든 값 반환
        while(iter.hasNext()){
            String str = iter.next();
            System.out.println(str);
        }
    }
}
