package com.fastcampus.programing.dmaker.repository;


import com.fastcampus.programing.dmaker.entity.Developer;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

/*
* DeveloperRepository
* - developer 엔티티를 통해 생성된 데이터를 H2데이터베이스에 넣는역할
* */


// repository Bean으로 등록
@Repository
public interface DeveloperRepository
        extends JpaRepository<Developer, Long> {
}
