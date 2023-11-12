package com.fastcampus.programing.dmaker.config;


import org.springframework.context.annotation.Configuration;
import org.springframework.data.jpa.repository.config.EnableJpaAuditing;

@Configuration // configuration 으로 등록하기 위함
@EnableJpaAuditing // CreatedDate, LastModifiedDate -> 두개의 JPA데이터 에디팅 어노테이션을 사용하기 위해 반.드.시 선언!!
/*
* @EnableJpaAuditing 을 configuration로 등록함으로써 테스트에도 JPA데이터에 접근하기 위함
* */
public class JpaConfig {
}
