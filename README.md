# FayGateway 反向代理网关
支持HTTP、TCP、GRPC代理

## 部署方式
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
