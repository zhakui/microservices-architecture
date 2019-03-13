package com.zhkui.core.recommendation.service;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;

@SpringBootApplication
@EnableDiscoveryClient
public class RecommendationService {
    public static void main(String[] args) {
        SpringApplication.run(RecommendationService.class, args);
    }
}