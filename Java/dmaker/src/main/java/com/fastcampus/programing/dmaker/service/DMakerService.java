package com.fastcampus.programing.dmaker.service;


import com.fastcampus.programing.dmaker.code.StatusCode;
import com.fastcampus.programing.dmaker.dto.CreateDeveloper;
import com.fastcampus.programing.dmaker.dto.DeveloperDetailDto;
import com.fastcampus.programing.dmaker.dto.DeveloperDto;
import com.fastcampus.programing.dmaker.dto.EditDeveloper;
import com.fastcampus.programing.dmaker.entity.Developer;
import com.fastcampus.programing.dmaker.entity.RetiredDeveloper;
import com.fastcampus.programing.dmaker.exception.DMakerException;
import com.fastcampus.programing.dmaker.repository.DeveloperRepository;
import com.fastcampus.programing.dmaker.repository.RetiredDeveloperRepository;
import com.fastcampus.programing.dmaker.type.DeveloperLevel;
import lombok.NonNull;
import org.springframework.transaction.annotation.*;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.stream.Collectors;

import static com.fastcampus.programing.dmaker.exception.DMakerErrorCode.*;


/*
 * DMakerService
 * - developer 엔티티 껍데기에 실제 데이터를 넣고 이를 DeveloperRepository에 넘긴다.
 * */
@Service // -> service Bean으로 등록
@RequiredArgsConstructor
public class DMakerService {
    private final DeveloperRepository developerRepository;
    private final RetiredDeveloperRepository retiredDeveloperRepository;

    // ACID
    // Atomic
    // Consistency
    // Isolation
    // Durability
    // [ 디벨로퍼 생성 ]
    @Transactional // -> Transactional 한 코드에서 반복되고있는 패턴을, 어노테이션 기반 포인트컷으로 인해 반복되는 코드인 advice를 자동으로 넣어준다 : AOP
    public CreateDeveloper.Response createDeveloper(CreateDeveloper.Request request) {
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
                .statusCode(StatusCode.EMPLOYED)
                .name(request.getName())
                .age(request.getAge())
                .build();

        developerRepository.save(developer);   // db 저장
        return CreateDeveloper.Response.fromEntity(developer); // 리턴값 반환;
    }

    // [ 디벨로퍼 리스트 출력 ]
    @Transactional(readOnly = true)
    public List<DeveloperDto> getAllEmployedDevelopers() {
        return developerRepository.findDevelopersByStatusCodeEquals(StatusCode.EMPLOYED)
                .stream()// 데이터가 많을수 있으니 stream데이터로 꺼내온다
                .map(DeveloperDto::fromEntity)// 꺼내온 Developer 엔티티 데이터를 DeveloperDto에 넣고 반환값을 만든다.
                .collect(Collectors.toList());// List로 collect한다.
    }

    // [ 디벨로퍼 한개 디테일 출력 ]
    @Transactional(readOnly = true)
    public DeveloperDetailDto getDeveloperDetail(String memberId) {
        return developerRepository.findByMemberId(memberId)
                .map(DeveloperDetailDto::fromEntity) // 꺼내온 Developer 엔티티 데이터를 DeveloperDetailDto에 넣고 반환값을 만든다.
                .orElseThrow(() -> new DMakerException(NO_DEVELOPER));// 만약 해당하는 값이 없으면 exception처리
    }

    // [ 디벨로퍼 정보 수정 ]
    @Transactional
    public DeveloperDetailDto editDeveloper(String memberId, EditDeveloper.Request request) {
        validateDeveloper(request.getDeveloperLevel(), request.getExperienceYears());

        Developer developer = developerRepository.findByMemberId(memberId).orElseThrow(() -> {
            throw new DMakerException(NO_DEVELOPER);
        });

        // @Transctional 어노테이션이 있다면 이미 저장된 데이터에대해서 수정후(setter)에 db에 save하지 않아도 된다.
        developer.setDeveloperLevel(request.getDeveloperLevel());
        developer.setDeveloperSkillType(request.getDeveloperSkillType());
        developer.setExperienceYears(request.getExperienceYears());


        return DeveloperDetailDto.fromEntity(developer);
    }

    // [ 디벨로퍼 정보 삭제 ]
    @Transactional
    public DeveloperDetailDto deleteDeveloper(String memberId) {
        // 1.EMPLOYED -> RETIRED
        Developer developer = developerRepository.findByMemberId(memberId)
                .orElseThrow(() -> new DMakerException(NO_DEVELOPER));
        developer.setStatusCode(StatusCode.RETIRED);

        // 2.save into RetiredDeveloper
        RetiredDeveloper retiredDeveloper = RetiredDeveloper.builder()
                .memberId(memberId)
                .name(developer.getName())
                .build();

        retiredDeveloperRepository.save(retiredDeveloper);
        return DeveloperDetailDto.fromEntity(developer);
    }




    // ============================================================ request business validation ============================================================
    // developer create validation
    private void validateCreateDeveloperRequest(@NonNull CreateDeveloper.Request request) {
        validateDeveloper(request.getDeveloperLevel(), request.getExperienceYears());

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



    private static void validateDeveloper(DeveloperLevel developerLevel, Integer experienceYears) {
        // 시니어 검사
        if (developerLevel == DeveloperLevel.SENIOR && experienceYears < 10) {
            throw new DMakerException(LEVEL_EXPERIENCE_YEARS_NOT_MATCHED);
        }

        // 중니어 검사
        if (developerLevel == DeveloperLevel.JUNGNIOR
                && (experienceYears < 4 || experienceYears > 10)) {
            throw new DMakerException(LEVEL_EXPERIENCE_YEARS_NOT_MATCHED);
        }

        // 주니어 검사
        if (developerLevel == DeveloperLevel.JUNIOR && experienceYears > 4) {
            throw new DMakerException(LEVEL_EXPERIENCE_YEARS_NOT_MATCHED);
        }
    }


}
