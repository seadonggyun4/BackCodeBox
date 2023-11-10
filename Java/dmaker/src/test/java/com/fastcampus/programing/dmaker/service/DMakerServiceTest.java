package com.fastcampus.programing.dmaker.service;

import com.fastcampus.programing.dmaker.dto.DeveloperDetailDto;
import com.fastcampus.programing.dmaker.entity.Developer;
import com.fastcampus.programing.dmaker.repository.DeveloperRepository;
import com.fastcampus.programing.dmaker.repository.RetiredDeveloperRepository;
import com.fastcampus.programing.dmaker.type.DeveloperLevel;
import com.fastcampus.programing.dmaker.type.DeveloperSkillType;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.BDDMockito;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.anyString;

@ExtendWith(MockitoExtension.class) // 데이터 모킹을 위한 Mockito사용을 위함
class DMakerServiceTest {
    @Mock
    private DeveloperRepository developerRepository;

    @Mock
    private RetiredDeveloperRepository retiredDeveloperRepository;

    @InjectMocks// -> 가짜 데이터를 주입
    private DMakerService dMakerService;


    @Test
    public void testSomething() {
        // Mocking: developerRepository 에 가짜 Mock데이터를 주입한다.
        BDDMockito.given(developerRepository.findByMemberId(anyString()))
                .willReturn(Optional.of(Developer.builder()
                        .developerLevel(DeveloperLevel.JUNIOR)
                        .developerSkillType(DeveloperSkillType.FRONT_END)
                        .experienceYears(2)
                        .memberId("testId")
                        .name("test")
                        .age(28)
                        .build()));

        DeveloperDetailDto developerDetail = dMakerService.getDeveloperDetail("memberId");
        System.out.println(developerDetail);

        assertEquals(DeveloperLevel.JUNIOR, developerDetail.getDeveloperLevel());
        assertEquals(DeveloperSkillType.FRONT_END, developerDetail.getDeveloperSkillType());
    }

}
