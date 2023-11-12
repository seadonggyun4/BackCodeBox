package com.fastcampus.programing.dmaker.service;

import com.fastcampus.programing.dmaker.code.StatusCode;
import com.fastcampus.programing.dmaker.dto.CreateDeveloper;
import com.fastcampus.programing.dmaker.dto.DeveloperDetailDto;
import com.fastcampus.programing.dmaker.entity.Developer;
import com.fastcampus.programing.dmaker.repository.DeveloperRepository;
import com.fastcampus.programing.dmaker.repository.RetiredDeveloperRepository;
import com.fastcampus.programing.dmaker.type.DeveloperLevel;
import com.fastcampus.programing.dmaker.type.DeveloperSkillType;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.ArgumentCaptor;
import org.mockito.BDDMockito;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.BDDMockito.given;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;

@ExtendWith(MockitoExtension.class) // 데이터 모킹을 위한 Mockito사용을 위함
class DMakerServiceTest {
    @Mock // 가짜 Repository에 접근하기 위함
    private DeveloperRepository developerRepository;

    @Mock
    private RetiredDeveloperRepository retiredDeveloperRepository;

    @InjectMocks// -> 가짜 데이터를 주입
    private DMakerService dMakerService;


    @Test
    public void getDeveloperDetailTest_SUCCESS() {
        //[ given ]
        // Mocking: developerRepository 에 가짜 Mock데이터를 주입한다.
        given(developerRepository.findByMemberId(anyString()))
                .willReturn(Optional.of(Developer.builder()
                        .developerLevel(DeveloperLevel.JUNIOR)
                        .developerSkillType(DeveloperSkillType.FRONT_END)
                        .experienceYears(2)
                        .statusCode(StatusCode.EMPLOYED)
                        .memberId("testId")
                        .name("test")
                        .age(28)
                        .build()));

        //[ when ]
        DeveloperDetailDto developerDetail = dMakerService.getDeveloperDetail("memberId");
        System.out.println(developerDetail);

        //[ then ]
        assertEquals(DeveloperLevel.JUNIOR, developerDetail.getDeveloperLevel());
        assertEquals(DeveloperSkillType.FRONT_END, developerDetail.getDeveloperSkillType());
    }

    @Test
    public void createDeveloperTest_SUCCESS() {
        //[ given ]
        CreateDeveloper.Request request =
                CreateDeveloper.Request.builder()
                        .developerLevel(DeveloperLevel.JUNIOR)
                        .developerSkillType(DeveloperSkillType.FRONT_END)
                        .experienceYears(2)
                        .memberId("testId")
                        .name("test")
                        .age(28)
                        .build();
        given(developerRepository.findByMemberId(anyString()))
                .willReturn(Optional.empty());
        ArgumentCaptor<Developer> captor = ArgumentCaptor.forClass(Developer.class); // Developer 엔티티에 들어오는 값의 캡쳐 객체 생성

        //[ when ]
        CreateDeveloper.Response response = dMakerService.createDeveloper(request);

        //[ then ]
        verify(developerRepository, times(1))
                .save(captor.capture()); // 가짜 developerRepository 에 captor객체가 넘어오는 순간을 캡쳐
        Developer savedDeveloper = captor.getValue(); // 캡쳐된 값을 변수에 할당

        // 캡처된 데이터에서 의도된결과값을 비교
        assertEquals(savedDeveloper.getDeveloperLevel(), DeveloperLevel.JUNIOR);
        assertEquals(savedDeveloper.getDeveloperSkillType(), DeveloperSkillType.FRONT_END);

    }

}
