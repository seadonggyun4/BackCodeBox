package org.example;

import org.example.logic.BubbleSort;
import org.example.logic.JavaSort;
import org.example.logic.Sort;
import org.example.service.SortService;

import java.util.Arrays;

public class Main {
    private static SortService sut = new SortService(new JavaSort<String>());


    public static void main(String[] args) {
        System.out.println( sut.doSort(Arrays.asList(args)));
    }
}