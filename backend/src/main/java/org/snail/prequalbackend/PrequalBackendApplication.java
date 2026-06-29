package org.snail.prequalbackend;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cache.annotation.EnableCaching;

@SpringBootApplication
@EnableCaching
public class PrequalBackendApplication {

    public static void main(String[] args) {
        SpringApplication.run(PrequalBackendApplication.class, args);
    }
}
