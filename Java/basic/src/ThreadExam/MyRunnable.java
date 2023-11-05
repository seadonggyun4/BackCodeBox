package ThreadExam;


// 1. Runnable 인터페이스를 구현한다.
public class MyRunnable implements Runnable{
    private String str;

    public MyRunnable(String str){
        this.str = str;
    }

    // 2. run()메서드를 오버라이딩
    @Override
    public void run(){
        for(int i=0; i < 10; i++){
            System.out.print(str);
            try {
                Thread.sleep(1000);
            } catch (InterruptedException e){
                e.printStackTrace();
            }
        }
    }
}
