# gin+viper+zap+gorm
```text
基础底座，没有业务代码
```

# 目录结构
```shell
conf               #配置文件目录
config             #配置文件viper
controller         #业务代码存放地
log                #程序日志文件存放
logger             #日志zap
model              #数据库序列化
router             #路由文件
sql                #数据库sql文件
tools              #链接各种中间件
util               #公共函数
main.go            #总入口
```


# 下载并安装Gin
```shell
go get -u github.com/gin-gonic/gin
```
# 换源
```cmd
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```
# 编译
```cmd
make - 格式化 Go 代码, 并编译生成二进制文件"
make build - 编译 Go 代码, 生成二进制文件"
make run - 直接运行 Go 代码"
make gotool - 运行 Go 工具 'fmt' and 'vet'"
### 1 目标平台的体系架构（386、amd64、arm）
set GOARCH=amd64
### 2 目标平台的操作系统（darwin、freebsd、linux、windows）
set GOOS=linux
### 3 编译 使用-o指定你要生成的文件名称，勿需指定可以去掉（参考：go build main.go）
go biuld -o serverName mian.go
```

# 编译linux下使用的包
```cmd
go env -w GOOS=linux
go env -w GOARCH=adm64
```

# 状态码
```text
200： 成功
204： 成功，但无内容返回。OPTIONS 请求时返回
301： 永久移动。新地址输入到 Location 头中。
304： 未修改
401： 未授权 请求要求身份验证。 对于需要登录的网页，服务器可能返回此响应。
403： 由于权限原因，拒绝访问。
404： 资源未找到。
405： 请求方法被禁止。
408： 请求超时，如上传大文件超时。
410： 接口已过期
500： 内部服务器错误
502： 网关错误
```