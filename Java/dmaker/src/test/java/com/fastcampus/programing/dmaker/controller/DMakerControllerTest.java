package com.fastcampus.programing.dmaker.controller;

import com.fastcampus.programing.dmaker.dto.DeveloperDto;
import com.fastcampus.programing.dmaker.service.DMakerService;
import com.fastcampus.programing.dmaker.type.DeveloperLevel;
import com.fastcampus.programing.dmaker.type.DeveloperSkillType;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.ResultMatcher;

import java.nio.charset.StandardCharsets;
import java.util.Arrays;

import static org.hamcrest.Matchers.is;
import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.BDDMockito.given;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;

import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(DMakerController.class)// 컨트롤러 테스트를 위함 -> 테스트할 컨트롤러Bean을 올린다.
class DMakerControllerTest {
    @Autowired
    private MockMvc mockMvc;

    @MockBean // -> 가짜Bean등록 : DMakerController 가 의존하고 있는 DMakerService
    private DMakerService dMakerService;

    protected MediaType contentType = new MediaType(
            MediaType.APPLICATION_JSON.getType(),
            MediaType.APPLICATION_JSON.getSubtype(),
            StandardCharsets.UTF_8
    ); // JSON형식의 UTF_8 타입 변수

    // [ /developers API 테스트 ]
    @Test
    void getAllDevelopersTest() throws Exception {
        DeveloperDto juniorDeveloperDto = DeveloperDto.builder()
                .developerLevel(DeveloperLevel.JUNIOR)
                .developerSkillType(DeveloperSkillType.FRONT_END)
                .memberId("testID1")
                .build();

        DeveloperDto seniorDeveloperDto = DeveloperDto.builder()
                .developerLevel(DeveloperLevel.SENIOR)
                .developerSkillType(DeveloperSkillType.FRONT_END)
                .memberId("testID2")
                .build();

        given(dMakerService.getAllEmployedDevelopers())
                .willReturn(Arrays.asList(juniorDeveloperDto, seniorDeveloperDto));

        // 가짜 mockMvc가 "/developers" 경로로 get테스트를 보낸다.
        mockMvc.perform(get("/developers")
                        .contentType(contentType))
                        .andExpect(status().isOk())
                        .andDo(print())
                        .andExpect(
                                jsonPath("$.[0].developerLevel", is(DeveloperLevel.JUNIOR.name())
                                ) //andExpect 내부의 jsonPath을 통해 해당 데이터를 비교하는 테스트 코드 작성 가능
                        );
    }
}