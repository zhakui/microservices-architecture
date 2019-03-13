#!/usr/bin/env bash

cd microservices/composite/productcomposite;     mvn spring-boot:run; cd -
cd microservices/core/product-service;           mvn spring-boot:run; cd -
cd microservices/core/recommendation-service;    mvn spring-boot:run; cd -
cd microservices/core/review-service;            mvn spring-boot:run; cd -

cd microservices/support/discovery-service;      mvn spring-boot:run; cd -