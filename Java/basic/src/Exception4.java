import java.io.FileInputStream;
import java.io.FileNotFoundException;

/*
* fileNotFoundException 은 "checked exception" 이다.
* checked exception은 컴파일 단계부터 예외처리를 하지 않으면 에러가 발생한다.
* */
public class Exception4 {
    public static void main(String[] args) {
        try {
            FileInputStream fils = new FileInputStream("Exception4.java");
        } catch (FileNotFoundException fnfe) {
            System.out.println(fnfe);
        }
    }
}
