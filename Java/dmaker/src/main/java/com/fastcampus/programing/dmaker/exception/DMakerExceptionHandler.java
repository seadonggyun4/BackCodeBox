package com.fastcampus.programing.dmaker.exception;


import com.fastcampus.programing.dmaker.dto.DMakerErrorResponse;
import jakarta.servlet.http.HttpServletRequest;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.HttpRequestMethodNotSupportedException;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;
@Slf4j
@RestControllerAdvice
// -> 각 컨트롤러에 조언을 해주는 어드바이스 역할을 하게 한다. Bean등록
// -> 전역적인 ExceptionHandler를 만들수 있다.
public class DMakerExceptionHandler {
    @ExceptionHandler(DMakerException.class) // -> HTTP통신시 DMakerException이 발생하면
    public DMakerErrorResponse handleException(DMakerException e, HttpServletRequest request) {
        log.error("errorCode: {},  url: {},  message: {}", e.getDMakerErrorCode(), request.getRequestURI(), e.getDetailMessage());
        return DMakerErrorResponse.builder()
                .errorCode(e.getDMakerErrorCode())
                .errorMessage(e.getDetailMessage())
                .build();
    }

    @ExceptionHandler(value = {
            HttpRequestMethodNotSupportedException.class, // 정확하지 않은 url
            MethodArgumentNotValidException.class // validation에 맞지않는 파라메터값
    })
    public DMakerErrorResponse handleBadRequest(
        Exception e, HttpServletRequest request
    ){
        log.error("url: {},  message: {}",  request.getRequestURI(), e.getMessage());
        return DMakerErrorResponse.builder()
                .errorCode(DMakerErrorCode.INVALID_REQUEST)
                .errorMessage(DMakerErrorCode.INVALID_REQUEST.getMessage())
                .build();
    }


    @ExceptionHandler(Exception.class)
    public DMakerErrorResponse handleException(
            Exception e, HttpServletRequest request
    ){
        log.error("url: {},  message: {}",  request.getRequestURI(), e.getMessage());
        return DMakerErrorResponse.builder()
                .errorCode(DMakerErrorCode.INVALID_REQUEST)
                .errorMessage(DMakerErrorCode.INVALID_REQUEST.getMessage())
                .build();
    }
}
