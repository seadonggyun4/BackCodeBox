//package Thread;
package ThreadExam;


public class MyThread extends Thread{
    private String str;


    public MyThread(String str){
        this.str = str;
    }


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
