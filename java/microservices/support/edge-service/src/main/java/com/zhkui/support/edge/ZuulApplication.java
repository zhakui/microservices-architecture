package com.zhkui.support.edge;

import org.springframework.boot.autoconfigure.SpringBootApplication;
<<<<<<< HEAD
import org.springframework.boot.SpringApplication;
import org.springframework.cloud.netflix.zuul.EnableZuulProxy;

@SpringBootApplication
=======
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.cloud.netflix.zuul.EnableZuulProxy;
import org.springframework.stereotype.Controller;

@SpringBootApplication
@Controller
>>>>>>> 1badb4e2d4fd7770aa200a15a31b7657a79f77a7
@EnableZuulProxy
public class ZuulApplication {
	
    public static void main(String[] args) {
<<<<<<< HEAD
        SpringApplication.run(ZuulApplication.class,args);
=======
        new SpringApplicationBuilder(ZuulApplication.class).web(true).run(args);
>>>>>>> 1badb4e2d4fd7770aa200a15a31b7657a79f77a7
    }
}
