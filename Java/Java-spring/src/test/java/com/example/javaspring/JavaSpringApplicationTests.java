package com.example.javaspring;

import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
class JavaSpringApplicationTests {

	@Test
	void contextLoads() {
		//Given
		String[] args = {"1", "3", "2"};

		//When & Then
		JavaSpringApplication.main(args);
	}

}
