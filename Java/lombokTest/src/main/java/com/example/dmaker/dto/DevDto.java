package com.example.dmaker.dto;

import lombok.*;

import java.time.LocalDateTime;

//@Getter
//@Setter
@ToString
@NoArgsConstructor
@AllArgsConstructor
@RequiredArgsConstructor
@Builder // 자바 생정자 선언 및 setter를 간략화 할수 있다.


// Data 는 위의 어노테이션을 전부 선언한것과 동일 -> 데이터 구조체를 위한 전용 어노테이션
// 불필요한 어노테이션이 오버라이딩 되는 경우도 있어 자주 사용 x
//@Data
public class DevDto {
    @NonNull
    String name;
    Integer age;
    LocalDateTime startAt;
}
