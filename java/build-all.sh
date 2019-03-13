#!/usr/bin/env bash

cd microservices/composite/productcomposite;     mvn clean; mvn package; cd -
cd microservices/core/product-service;           mvn clean; mvn package; cd -
cd microservices/core/recommendation-service;    mvn clean; mvn package; cd -
cd microservices/core/review-service;            mvn clean; mvn package; cd -

cd microservices/support/discovery-service;      mvn clean; mvn package; cd -
