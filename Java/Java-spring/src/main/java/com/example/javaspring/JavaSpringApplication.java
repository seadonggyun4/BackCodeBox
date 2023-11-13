package com.example.javaspring;

import com.example.javaspring.config.Config;
import com.example.javaspring.logic.JavaSort;
import com.example.javaspring.logic.Sort;
import com.example.javaspring.service.SortService;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;

import java.util.Arrays;

@SpringBootApplication
public class JavaSpringApplication {


	public static void main(String[] args) {
		ApplicationContext context = new AnnotationConfigApplicationContext(Config.class);
		SortService sortService = context.getBean(SortService.class);

		 System.out.println( sortService.doSort(Arrays.asList(args)));
	}

}
