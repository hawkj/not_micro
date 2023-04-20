#not_micro框架的定位：
1. 适用于初创型公司，后端技术人数小于20人。
2. not_micro 即不使用微服务技术。

#可提供的能力
1. 程序的链路追踪，基于jaeger实现。
2. 服务限流。
3. JWT等中间件。
4. 方便的异步消费JOB。
5. CI/DI 友善。支持多服务分开部署。
6. 统一的配置中心，基于nacos。

#quick start
## 环境搭建
### 在服务器上创建文件夹用来保存容器中的持久花数据
1. mkdir -p -m 777 /var/container_data/
### 启动环境
1. cd <not_micro根目录>
2. docker-compose up -d



