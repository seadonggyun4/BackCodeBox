import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class SortShuffleExam {
    public static void main(String[] args) {
        List<String> list = new ArrayList<>();
        list.add("kim");
        list.add("lee");
        list.add("hong");

        Collections.sort(list); // Collections 내장클래스의 sort매서드에 객체변수를 넣어주면 오름차순으로 자동정렬이 된다.

        // 생성자 타입중 Comparable이라 불리는 타입을 가지고있는 객체만 정렬이 가능하다.

//        for(int i = 0;  i < list.size(); i++){
//            System.out.println(list.get(i));
//        }

        for (String i : list){
            System.out.println(i);
        }



        List<String> list2 = new ArrayList<>();
        list2.add("kim");
        list2.add("lee");
        list2.add("hong");

        Collections.shuffle(list2);// Collections 내장클래스의 shuffle매서드에 객체변수를 넣어주면 랜덤하게 섞인다..

        for (String i : list2){
            System.out.println(i);
        }
    }
}
