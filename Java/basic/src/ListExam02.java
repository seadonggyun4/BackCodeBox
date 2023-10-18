import java.util.ArrayList;
import java.util.List;

public class ListExam02 {
    public static void main(String[] args){
        // generic을 사용하여 리스트를 String타입으로 선언
        List<String> list = new ArrayList<>();
        list.add("kim");
        list.add("lee");
        list.add("hong");


        // Object 타입을 String으로 형 변환이 필요없다.
        String str1 = list.get(0);
        String str2 = list.get(1);
        String str3 = list.get(2);

        System.out.println(str1);
        System.out.println(str2);
        System.out.println(str3);
    }
}
