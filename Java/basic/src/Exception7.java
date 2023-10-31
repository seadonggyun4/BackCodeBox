public class Exception7 {
    public static void main(String[] args) {
        try {
            ExceptionObj7 exobj7 = new ExceptionObj7();
            int value = exobj7.divide(10 ,0);
            System.out.println(value);
        } catch (MyException ex){
            System.out.println(ex);
            System.out.println("사용자 정의 Exception이 발생했네요!");
        }
    }
}


class ExceptionObj7 {
    public int divide(int i, int k) throws MyException {
        int value = 0;
        try {
            value = i / k;
        } catch (ArithmeticException ae){
            throw new MyException("0으로 나눌 수 없습니다.");
        }
        return value;
    }
}
