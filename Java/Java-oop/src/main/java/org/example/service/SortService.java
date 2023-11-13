package org.example.service;

import org.example.logic.JavaSort;
import org.example.logic.Sort;

import java.util.List;

public class SortService {
    // 정렬구현체 생성
    private final Sort<String> sort;

    // 어떤 정렬 구현체를 받을것인지 매개변수로 받는다.
    public SortService(Sort<String> sort){
        this.sort = sort;
    }

    public List<String> doSort(List<String> list) {
        return sort.sort(list);
    }
}
