package ThreadExam;

public class MyThreadExam {
    public static void main(String[] args){
        MyThread mt1 = new MyThread("*");
        MyThread mt2 = new MyThread("+");

        // 3. thread는 start()메서드로 실행
        mt1.start();
        mt2.start();
    }
}
