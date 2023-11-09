package com.fastcampus.programing.dmaker.controller;


import com.fastcampus.programing.dmaker.dto.CreateDeveloper;
import com.fastcampus.programing.dmaker.service.DMakerService;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
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


    // developers 리스트 출력 API ========================================
    @GetMapping("/developers") // -> HTTP GET요청 매핑
    public List<String> getAllDevelopers() {
        log.info("GET /developers HTTP/1.1");

        return Arrays.asList("snow", "elsa", "oulaf");
    }

    // developers 생성 API =============================================
    @PostMapping("/create-developers") // -> HTTP GET요청 매핑
    public CreateDeveloper.Response createDevelopers(
            @Valid // -> @Valid어노테이션을 통해 CreatteDeveloper내무에 선언한 @Min, @Max 와 같은 validation 어노테이션이 동작한다.
            @RequestBody // -> HTTP스팩의 json을 담는다.
            CreateDeveloper.Request request
            ) {
        log.info("request: {}", request);

        dMakerService.createDeveloper(request);

        return dMakerService.createDeveloper(request);
    }
}
