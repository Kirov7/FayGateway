# FayGateway 反向代理网关
支持HTTP、TCP、GRPC代理

部署方式
```shell
go mod tidy
go run main.go -conf ./conf/dev/ -endpoint dashboard
```
另起一个进程
```shell
go run main.go -conf ./conf/dev/ -endpoint server
```
