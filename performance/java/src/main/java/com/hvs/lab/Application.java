package com.hvs.lab;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;

@RestController
class Application {

    @GetMapping("ping")
    public String ping(){
        for (int i = 0; i < 1000000; i++) {
            System.out.println("loop count: " + i);
        }
        return "pong";
    }

}