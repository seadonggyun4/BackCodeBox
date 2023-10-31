/*
* ArithmeticException 은 "runtime exception" 이다.
* 컴파일 이 아닌 실행도중에 에러가 발생한다.
* */
public class Exception2 {
    public static void main(String[] args) {
        ExceptionObj2 exobj = new ExceptionObj2();
        try {
            int value = exobj.divide(10, 0);
            System.out.println(value);
        } catch (ArithmeticException ex) {
            System.out.println("0으로 숫자를 나눌수 없습니다.");
            System.out.println(ex.toString());
        }

    }
}


