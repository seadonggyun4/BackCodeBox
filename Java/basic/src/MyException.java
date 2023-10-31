public class MyException extends RuntimeException{
    // 오류 메시지나, 발생한 Exception을 감싼 결과로 내가 만든 Exception을 사용하고 싶을 때가 많다.

    public MyException(String message) {
        super(message); // MyException생성자는 부모에게 message전달
    }

    public MyException(Exception cause){
        super(cause);
    }
}
