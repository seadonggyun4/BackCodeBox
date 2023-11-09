package com.fastcampus.programing.dmaker.service;


import com.fastcampus.programing.dmaker.dto.CreateDeveloper;
import com.fastcampus.programing.dmaker.entity.Developer;
import com.fastcampus.programing.dmaker.exception.DMakerErrorCode;
import com.fastcampus.programing.dmaker.exception.DMakerException;
import com.fastcampus.programing.dmaker.repository.DeveloperRepository;
import com.fastcampus.programing.dmaker.type.DeveloperLevel;
import com.fastcampus.programing.dmaker.type.DeveloperSkillType;
import jakarta.persistence.EntityManager;
import jakarta.transaction.Transactional;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.Optional;

import static com.fastcampus.programing.dmaker.exception.DMakerErrorCode.*;


/*
* DMakerService
* - developer 엔티티 껍데기에 실제 데이터를 넣고 이를 DeveloperRepository에 넘긴다.
* */
@Service // -> service Bean으로 등록
@RequiredArgsConstructor
public class DMakerService {
    private final DeveloperRepository developerRepository;
    private final EntityManager em;




    // ACID
    // Atomic
    // Consistency
    // Isolation
    // Durability
    @Transactional // -> Transactional 한 코드에서 반복되고있는 패턴을, 어노테이션 기반 포인트컷으로 인해 반복되는 코드인 advice를 자동으로 넣어준다 : AOP
    public CreateDeveloper.Response createDeveloper(CreateDeveloper.Request request){
        /*
        * 여러 작업을 진행하다가 문제가 생겼을 경우 이전 상태로 롤백하기 위해 사용되는 것이 트랜잭션(Transaction) 이다.
        * 트랜잭션은 더 이상 쪼갤 수 없는 최소 작업 단위를 의미한다. 그래서 트랜잭션은 commit으로 성공 하거나 rollback으로 실패 이후 취소되어야 한다.
        * 하지만 모든 트랜잭션이 동일한 것은 아니고 속성에 따라 동작 방식을 다르게 해줄 수 있다.
        *
        *
        * => 동시에 진행되고 성공해야 하는 작업중 일부가 실패했을때 다시 전체적으로 그 전 작업으로 롤백하기 위함
        * */


        // ==================== @Transactional 어노테이션 X ====================
//        EntityTransaction transction = em.getTransaction();
//        try{
//            transction.begin();
//            Developer developer = Developer.builder()
//                    .developerLevel(DeveloperLevel.JUNGNIOR)
//                    .developerSkillType(DeveloperSkillType.FRONT_END)
//                    .experienceYears(2)
//                    .name("Olaf")
//                    .age(5)
//                    .build();
//
//            developerRepository.save(developer);
//            transction.commit();
//        } catch (Exception e) {
//            transction.rollback();
//            throw e;
//        }

        // ==================== @Transactional 어노테이션 O ====================
        validateCreateDeveloperRequest(request);


        Developer developer = Developer.builder()
                .developerLevel(request.getDeveloperLevel())
                .developerSkillType(request.getDeveloperSkillType())
                .experienceYears(request.getExperienceYears())
                .memberId(request.getMemberId())
                .name(request.getName())
                .age(request.getAge())
                .build();

        developerRepository.save(developer);   // db 저장
        return CreateDeveloper.Response.fromEntity(developer); // 리턴값 반환;
    }


    // ============================================================ request business validation ============================================================
    private void validateCreateDeveloperRequest(CreateDeveloper.Request request){
        DeveloperLevel developerLevel = request.getDeveloperLevel();
        Integer experienceYears = request.getExperienceYears();

        // 시니어 검사
        if(developerLevel == DeveloperLevel.SENIOR && experienceYears < 10) {
            throw new DMakerException(LEVEL_EXPERIENCE_YEARS_NOT_MATCHED);
        }

        // 중니어 검사
        if(developerLevel == DeveloperLevel.JUNGNIOR
                && (experienceYears < 4 || experienceYears > 10)) {
            throw new DMakerException(LEVEL_EXPERIENCE_YEARS_NOT_MATCHED);
        }

        // 주니어 검사
        if(developerLevel == DeveloperLevel.JUNIOR && experienceYears > 4) {
            throw new DMakerException(LEVEL_EXPERIENCE_YEARS_NOT_MATCHED);
        }

        // 멤버 아이디 존재여부 검사
        /*
        * Optional을 사용하면 예상치 못한 NullPointerException 예외를 제공되는 메소드로 간단히 회피할 수 있다.
        * 즉, 복잡한 조건문 없이도 널(null) 값으로 인해 발생하는 예외를 처리할 수 있게 된다.
        *
        * - isPresent() 메소드 : Optional 객체가 값을 가지고 있다면 true, 값이 없다면 false 리턴
        * */
//        Optional<Developer> developer = developerRepository.findByMemberId(request.getMemberId());
//        if(developer.isPresent()){
//            throw new DMakerException(DUPLICATED_MEMBER_ID);
//        }
        developerRepository.findByMemberId(request.getMemberId()).ifPresent(developer -> {
            throw new DMakerException(DUPLICATED_MEMBER_ID);
        });
    }
}
