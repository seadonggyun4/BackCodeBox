package ThreadExam;

public class MyThreadExam2 {
    public static void main(String[] args){
        MyRunnable mr1 = new MyRunnable("*");
        MyRunnable mr2 = new MyRunnable("+");

        // 3. thread인스턴스를 생성하는데, 생성자에 Runnable인스턴스를 넣어준다.
        Thread t1 = new Thread(mr1);
        Thread t2 = new Thread(mr2);

        // 4. Thread가 가지고있는 start매서드 호출
        t1.start();
        t2.start();

    }
}
