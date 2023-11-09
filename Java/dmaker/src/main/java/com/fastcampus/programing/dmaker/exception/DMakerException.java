package com.fastcampus.programing.dmaker.exception;

import lombok.Getter;

@Getter
public class DMakerException extends RuntimeException{
    private DMakerErrorCode dMakerErrorCode;
    private String detailMessage;

    // DMakerException 의 매개변수가 없을때
    // 상속받은 부모 RuntimeException 의 메시지를 넣는다.
    public DMakerException(DMakerErrorCode errorCode){
        super(errorCode.getMessage()); // runtimeExpection에서 에러메시지를 받아온다.
        this.dMakerErrorCode = errorCode;
        this.detailMessage = errorCode.getMessage();
    }


    // DMakerException 의 매개변수가 있을때
    public DMakerException(DMakerErrorCode errorCode, String detailMessage){
        super(detailMessage); // -> RuntimeException의 생성자에 매개변수 메시지를 넣어 RuntimeException의 에러베시지도 할당해준다.
        this.dMakerErrorCode = errorCode;
        this.detailMessage = detailMessage;
    }
}
