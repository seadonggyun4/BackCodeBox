package com.fastcampus.programming.dmaker.service;


import com.fastcampus.programming.dmaker.entity.Developer;
import com.fastcampus.programming.dmaker.repository.DeveloperRepository;
import com.fastcampus.programming.dmaker.type.DeveloperLevel;
import com.fastcampus.programming.dmaker.type.DeveloperSkillType;
import jakarta.transaction.Transactional;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;


/*
* DMakerService
* - developer 엔티티 껍데기에 실제 데이터를 넣고 이를 DeveloperRepository에 넘긴다.
* */
@Service // -> service Bean으로 등록
@RequiredArgsConstructor
public class DMakerService {
    private final DeveloperRepository developerRepository;


    // ACID
    // Atomic
    // Consistency
    // Isolation
    // Durability
    @Transactional
    public void createDeveloper(){
        Developer developer = Developer.builder()
                .developerLevel(DeveloperLevel.JUNGNIOR)
                .developerSkillType(DeveloperSkillType.FRONT_END)
                .experienceYears(2)
                .name("Olaf")
                .age(5)
                .build();

        developerRepository.save(developer);
    }
}
