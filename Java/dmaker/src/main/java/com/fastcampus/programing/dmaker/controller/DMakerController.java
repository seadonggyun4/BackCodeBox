package com.fastcampus.programing.dmaker.controller;


import com.fastcampus.programing.dmaker.service.DMakerService;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Arrays;
import java.util.Collections;
import java.util.List;


/*
* DMakerController
* - 각 HTTP에 대한 정의와 경로 매핑
*
* */

@Slf4j // lombok 어노테이션 -> 시스템 log를 찍을수 있게 한다.
@RestController // 컨트롤러 Bean 등록 ->  어노테이션 Controller + ResponseBody:  RestAPI선언 및 Json형식의 반환값
@RequiredArgsConstructor
public class DMakerController {
    private final DMakerService dMakerService;

    @GetMapping("/developers") // -> HTTP GET요청 매핑
    public List<String> getAllDevelopers(){
        log.info("GET /developers HTTP/1.1");

        return Arrays.asList("snow", "elsa", "oulaf");
    }


    @GetMapping("/create-developers") // -> HTTP GET요청 매핑
    public List<String> createDevelopers(){
        log.info("GET /create-developers HTTP/1.1");

        dMakerService.createDeveloper();

        return Collections.singletonList("Olaf");
    }
}
