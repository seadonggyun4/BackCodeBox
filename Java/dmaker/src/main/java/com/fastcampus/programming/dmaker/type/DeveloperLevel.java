package com.fastcampus.programming.dmaker.type;


import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor // -> 생성자를 만들어준다.
@Getter
public enum DeveloperLevel {
    New("신입 개발자"),
    JUNIOR("주니어 개발자"),
    JUNGNIOR("중니어 개발자"),
    SENIOR("시니어 개발자");

    private final String description;
}
