# FayGateway 反向代理网关
支持HTTP、TCP、GRPC的反向代理与负载均衡

## 部署方式
加载依赖
```shell
go mod tidy
```
启动控制面服务
```shell
go run main.go -conf ./conf/dev/ -endpoint dashboard
```
启动代理服务器服务
```shell
go run main.go -conf ./conf/dev/ -endpoint server
```
