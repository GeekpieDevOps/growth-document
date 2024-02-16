# Backend

## 运行指令

1. 创建目录，初始化模块：
 
```bash
cd backend
go mod init  github.com/GeekpieDevOps/growth-document
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

2.创建main.go文件，进行Hello World实例测试，以及数据库访问和端口连接
- 填写部分[内容](./main.go)(文件目前只写了一部分，含有Bug，具体见[文件注释](./main.go))

3.选择测试框架（待定testing），具体再商榷
- 编写[测试文件](./main_test.go)(空白)

## 运行测试方法指令







