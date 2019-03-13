package com.zhkui.core.review.service;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;

@SpringBootApplication
@EnableDiscoveryClient
public class ReviewService {
    public static void main(String[] args) {
        SpringApplication.run(ReviewService.class, args);
    }
}