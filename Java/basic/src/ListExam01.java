import java.util.ArrayList;

public class ListExam01 {
    public static void main(String[] args){
        // 자료구조 객체들은 제네릭을 사용하지 않으면
        // object타입을 저장합니다.
        ArrayList list = new ArrayList();
        list.add("kim");
        list.add("lee");
        list.add("hong");


        // Object 타입을 String으로 형 변환 해야한다.
        String str1 = (String)list.get(0);
        String str2 = (String)list.get(1);
        String str3 = (String)list.get(2);

        System.out.println(str1);
        System.out.println(str2);
        System.out.println(str3);
    }
}
