package com.hvs.lab;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

class MainTest {

    @Test
    void sayHello() {
        // Given
        String expectedMessage = "Hello world!";

        // When
        String message = Main.sayHello();

        // Then
        Assertions.assertEquals(expectedMessage, message);
    }
}