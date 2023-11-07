package com.example.dmaker.dto;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

import java.time.LocalDateTime;


class DevDtoTest {
    @Test
    void test(){
          DevDto devDto = DevDto.builder().name("snow").age(21).startAt(LocalDateTime.now()).build();
//          devDto.setName("snow");
//          devDto.setAge(21);
//          devDto.setStartAt(LocalDateTime.now());
          System.out.println(devDto);
    }



}