package org.example;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class MainTest {
    @Test
    void main(){
        //Given
        String[] args = {"1", "3", "2"};

        //When & Then
        Main.main(args);
    }
}