package com.fastcampus.programming.dmaker.entity;

import com.fastcampus.programming.dmaker.type.DeveloperLevel;
import com.fastcampus.programming.dmaker.type.DeveloperSkillType;
import jakarta.persistence.*;
import lombok.*;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import java.time.LocalDateTime;

/*
* Developer 테이블 객체
* - developer 에 필요한 데이터 컬럼과 타입을 선언
*
* */


@Getter
@Setter
@Builder
@NoArgsConstructor
@AllArgsConstructor
@Entity // -> Entity 선언
@EntityListeners(AuditingEntityListener.class)// auditing기능을 위한 어노테이션
public class Developer {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    protected  Long id;

    @Enumerated(EnumType.STRING)
    private DeveloperLevel developerLevel;


    @Enumerated(EnumType.STRING)
    private DeveloperSkillType developerSkillType;


    private Integer experienceYears;
    private String memberId;
    private String name;
    private Integer age;

    @CreatedDate // -> 자동으로 데이터 처음 세팅된 시간 기록
    private LocalDateTime createdAt;

    @LastModifiedDate // -> 자동으로 데이터 수정된 시간 기록
    private LocalDateTime updatedAt;
}
