package com.zhkui.composite.product;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
<<<<<<< HEAD
import org.springframework.cloud.client.circuitbreaker.EnableCircuitBreaker;
=======
>>>>>>> 1badb4e2d4fd7770aa200a15a31b7657a79f77a7
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;

@SpringBootApplication
@EnableDiscoveryClient
<<<<<<< HEAD
@EnableCircuitBreaker
=======
>>>>>>> 1badb4e2d4fd7770aa200a15a31b7657a79f77a7
public class ProductCompositeServiceApplication {
    public static void main(String[] args) {
        SpringApplication.run(ProductCompositeServiceApplication.class, args);
    }
}