public class Exception6 {
    public static void main(String[] args) {
        int[] array = {4, 2};
        int[] value = null;
        //        int[] value = new int[1];


        /*
        * 어떤 Exception 인지에 따라 다르게 핸들링 하기위해 어러번의 catch를 쓰고있다.
        * ArrayIndexOutOfBoundsException, ArithmeticException 모두 Exception의 하위 클래스이기때문에 먼저 실행후 두개의 익셉션에서
        * 걸리지 않은 에러가 Exception에서 실행되게 된다.
        * */
        try {
            value[0] = array[0] / array[1];
        } catch (ArrayIndexOutOfBoundsException  aiob){
             System.out.println(aiob.toString());
        } catch (ArithmeticException ae) {
            System.out.println(ae.toString());
        } catch (Exception ex) {
            System.out.println(ex);
        }
    }
}
