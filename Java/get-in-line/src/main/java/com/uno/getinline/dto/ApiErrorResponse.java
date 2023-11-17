package com.uno.getinline.dto;

import com.uno.getinline.constant.ErrorCode;
import lombok.*;


// 서비스 공통적으로 발생하는 에러 핸들
@Getter
@ToString
@EqualsAndHashCode
@RequiredArgsConstructor(access = AccessLevel.PROTECTED)
public class ApiErrorResponse {

    private final Boolean success;
    private final Integer errorCode;
    private final String message;


    // 생성자 매서드 =============================================================================================
    public static ApiErrorResponse of(Boolean success, Integer errorCode, String message) {
        return new ApiErrorResponse(success, errorCode, message);
    }

    public static ApiErrorResponse of(Boolean success, ErrorCode errorCode) {
        return new ApiErrorResponse(success, errorCode.getCode(), errorCode.getMessage());
    }

    public static ApiErrorResponse of(Boolean success, ErrorCode errorCode, Exception e) {
        return new ApiErrorResponse(success, errorCode.getCode(), errorCode.getMessage(e));
    }

    public static ApiErrorResponse of(Boolean success, ErrorCode errorCode, String message) {
        return new ApiErrorResponse(success, errorCode.getCode(), errorCode.getMessage(message));
    }

}
