# 基于哪个镜像
FROM java:8

VOLUME /tmp


# 拷贝文件到容器，也可以直接写成ADD microservice-discovery-eureka-0.0.1-SNAPSHOT.jar /app.jar
ADD target/product-service-1.0-SNAPSHOT.jar app.jar
RUN bash -c 'touch /app.jar'


# 开放8000端口
EXPOSE 8080


# 配置容器启动后执行的命令
ENTRYPOINT ["java","-Dspring.profiles.active=docker","-Djava.security.egd=file:/dev/./urandom","-jar","/app.jar"]

