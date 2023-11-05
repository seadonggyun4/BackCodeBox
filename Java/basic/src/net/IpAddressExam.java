package net;

import java.net.InetAddress;
import java.net.UnknownHostException;

public class IpAddressExam {
    public static void main(String[] args) {
        // ip4 형식 : 142.250.76.132
        // ip6 형식 : 2404:6800:400a:813:0:0:0:2004


        try {
            InetAddress ia = InetAddress.getLocalHost(); // 내컴퓨터 IP정보를 받아온다.
            System.out.println(ia.getHostAddress());
        } catch (UnknownHostException err){
            err.printStackTrace();
        }


        try {
            System.out.println("===============================================");
            InetAddress[] iaArr = InetAddress.getAllByName("www.google.com"); // google도메인의 IP정보를 리스트 형태로 받아온다.
            for(InetAddress ia : iaArr){
                System.out.println(ia.getHostAddress());
            }
        } catch (UnknownHostException err){
            err.printStackTrace();
        }


    }
}
