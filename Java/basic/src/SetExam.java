import java.util.*;
public class SetExam {
    public static void main(String[] args) {
        Set<String> set = new HashSet<>();
        set.add("hello");
        set.add("hi");
        set.add("hong");

        Iterator<String> iter = set.iterator();



        while (iter.hasNext()){
            String str = iter.next();
            System.out.println(str);
        }
    }
}
