package net;

import java.io.*;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.ArrayList;
import java.util.List;

public class VerySimpleWebServer {
    public static void main(String[] args) throws Exception{
        ServerSocket ss = new ServerSocket(9090);
        System.out.println("클라이언트 접속 대기");

        // 클라이언트의 연결을 대기
        // 클라이언트가 접속하는 순간, 클라이언트와 통신할 수 있는 socket객체 반환.
        Socket socket = ss.accept();

        // client와 읽고쓸수 있는 인풋,아웃풋 스트림이 생성된다.
        OutputStream out = socket.getOutputStream();
        InputStream in = socket.getInputStream();

//        byte[] buffer = new byte[512];
//        int readCount = 0;
//
//        while((readCount = in.read(buffer)) != -1){
//            System.out.write(buffer, 0, readCount);
//        }

        PrintWriter pw = new PrintWriter(new OutputStreamWriter(out));
        BufferedReader br = new BufferedReader(new InputStreamReader(in));
        String firstLine = br.readLine();
        List<String> headers = new ArrayList<>();
        String line = null;

        while (!(line = br.readLine()).equals("")){
            headers.add(line);
        }

        // 요청정보 읽기 끝
        System.out.println(firstLine);
        for(int i = 0; i < headers.size(); i++){
            System.out.println(headers.get(i));
        }

        pw.println("HTTP/ 200 OK");
        pw.println("name: seo");
        pw.println("email: rama4251@naver.com");
        pw.println();
        pw.println("<html>");
        pw.println("<h1>HELLO JAVA!!!!</h1>");
        pw.println("</html>");
        pw.close();

        ss.close();
        System.out.println("서버 종료");
    }
}
