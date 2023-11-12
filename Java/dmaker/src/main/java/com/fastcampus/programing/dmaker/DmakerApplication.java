package com.fastcampus.programing.dmaker;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.jpa.repository.config.EnableJpaAuditing;

//@EnableJpaAuditing // CreatedDate, LastModifiedDate -> 두개의 JPA데이터 에디팅 어노테이션을 사용하기 위해 반.드.시 선언!!
@SpringBootApplication
public class DmakerApplication {

    public static void main(String[] args) {
        SpringApplication.run(DmakerApplication.class, args);
    }

}
