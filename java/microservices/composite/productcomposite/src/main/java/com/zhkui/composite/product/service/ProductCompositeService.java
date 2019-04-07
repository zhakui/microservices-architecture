package com.zhkui.composite.product.service;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;

@SpringBootApplication
@EnableDiscoveryClient
public class ProductCompositeService {
    public static void main(String[] args) {
        SpringApplication.run(ProductCompositeService.class, args);
    }
}