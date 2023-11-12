package com.fastcampus.programing.dmaker.controller;


import com.fastcampus.programing.dmaker.dto.*;
import com.fastcampus.programing.dmaker.exception.DMakerException;
import com.fastcampus.programing.dmaker.service.DMakerService;

import jakarta.servlet.http.HttpServletRequest;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

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
    public List<DeveloperDto> getAllDevelopers() {
        log.info("GET /developers HTTP/1.1");

        return dMakerService.getAllEmployedDevelopers();
    }

    @GetMapping("/developer/{memberId}") // -> HTTP GET요청 매핑
    public DeveloperDetailDto getDeveloperDetail(
            @PathVariable final String memberId
    ) {
        log.info("GET /developer/{memberId} HTTP/1.1");

        return dMakerService.getDeveloperDetail(memberId);
    }

    // developers 생성 API =============================================
    @PostMapping("/create-developers") // -> HTTP POST요청 매핑
    public CreateDeveloper.Response createDevelopers(
            @Valid // -> @Valid어노테이션을 통해 CreatteDeveloper내무에 선언한 @Min, @Max 와 같은 validation 어노테이션이 동작한다.
            @RequestBody // -> HTTP스팩의 json을 담는다.
            final CreateDeveloper.Request  request
            ) {
        log.info("request: {}", request);

        return dMakerService.createDeveloper(request);
    }


    // developers 수정 API =============================================
    @PutMapping("/developer/{memberId}") // -> HTTP PUT요청 매핑
    public DeveloperDetailDto editDeveloper(
            @PathVariable final String memberId,
            @Valid
            @RequestBody
            final EditDeveloper.Request request
    ) {
        log.info("PUT //developer/{memberId} HTTP/1.1");

        return dMakerService.editDeveloper(memberId, request);
    }


    // developers 삭제 API =============================================
    @DeleteMapping("/developer/{memberId}") // -> HTTP DELETE요청 매핑
    public DeveloperDetailDto deleteDeveloper(
            @PathVariable final String memberId
    ) {
        log.info("DELETE /developer/{memberId} HTTP/1.1");

        return dMakerService.deleteDeveloper(memberId);
    }


    // Controller 예외처리 =============================================
    /*
     *  -> @ExceptionHandler 를 통한 예외 처리기법은 각 API내에서 try{} catch{}를 통해 예외처라하는 것보단 글로벌하다
     *  -> 하지만 각 컨트롤러 파일별로 @ExceptionHandler 를 계속 만들어줘야 하기때문에 이는 보일러 플레이트나 마찬가지 이다.
     * */


    //    @ResponseStatus(value = HttpStatus.CONFLICT)// 409
    //    @ExceptionHandler(DMakerException.class) // -> HTTP통신시 DMakerException이 발생하면
    //    public DMakerErrorResponse handleException(DMakerException e, HttpServletRequest request) {
    //        log.error("errorCode: {},  url: {},  message: {}", e.getDMakerErrorCode(), request.getRequestURI(), e.getDetailMessage());
    //        return DMakerErrorResponse.builder()
    //                .errorCode(e.getDMakerErrorCode())
    //                .errorMessage(e.getDetailMessage())
    //                .build();
    //    }
}
