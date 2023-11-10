package com.fastcampus.programing.dmaker.dto;

import com.fastcampus.programing.dmaker.entity.Developer;
import com.fastcampus.programing.dmaker.type.DeveloperLevel;
import com.fastcampus.programing.dmaker.type.DeveloperSkillType;
import lombok.*;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
@Builder
@ToString
public class DeveloperDto {
    private DeveloperLevel developerLevel;
    private DeveloperSkillType developerSkillType;
    private String memberId;

    public static DeveloperDto fromEntity(Developer developer){
        return DeveloperDto.builder()
                .developerLevel(developer.getDeveloperLevel())
                .developerSkillType(developer.getDeveloperSkillType())
                .memberId(developer.getMemberId())
                .build();
    }
}
