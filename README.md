


### 项目演示

商城后台演示：https://www.zimall.site

小程序商城演示：暂不支持，因个人主体小程序未开放电商相关类目，审核未通过😭

### 技术选型

| 技术 | 说明 | 相关文档 |
|---|---|---|
| vue3 | 前端框架 | https://v3.cn.vuejs.org |
| vue-router | 页面路由 | https://next.router.vuejs.org |
| axios | 网络请求库 | https://axios-http.com |
| vuex | 状态管理 | https://next.vuex.vuejs.org |
| element plus | 前端UI组件库 | https://element-plus.org |
| vant weapp | 微信小程序UI组件库 | https://vant-contrib.gitee.io/vant-weapp |
| gin | Web框架 | https://gin-gonic.com |
| gorm | ORM框架 | https://gorm.io |
| jwt | 用户认证 | https://github.com/golang-jwt/jwt |
| captcha | 验证码生成器 | https://github.com/mojocn/base64Captcha |
| viper | 配置管理 | https://github.com/spf13/viper |
| redis | 数据缓存 | https://github.com/go-redis/redis |

### 项目结构
```
imall
  ├── app         // 微信小程序
  ├── demo        // 演示资源
  ├── server      // 服务端
  ├── web         // 商家后台
  ├── ...         // 其他
```
### 开发工具

本项目使用 Visual Studio Code、Navicat Premium、微信开发者工具等开发工具。

### 本地运行

运行环境：

| 环境 | 版本 | 下载地址 |
|---|---|---|
| go | >= 1.17.1 | https://golang.google.cn/dl/ |
| mysql | >= 8.0.28 | https://www.mysql.com/downloads/ |
| redis | >= 6.0.16 | https://redis.io/download/ |
| node | >= 14.13.1 | https://nodejs.org/en/download/ |

直接下载压缩包，或使用 Git 克隆项目：
```
$ git clone https://github.com/zchengo/imall.git
```

**部署一：Go服务端（server）**

修改配置文件：配置文件位于 /server/config.yaml，请按实际情况进行修改。

推荐使用 Goland 或 VSCode 打开 server 目录，在 Terminal(终端) 中，执行如下命令。
```
$ cd server
$ go mod tidy
$ go build -o server main.go (windows编译命令为 go build -o server.exe main.go )

# 运行二进制
$ ./server (windows运行命令为 server.exe)
```

**部署二：商家后台（web）**

推荐使用 WebStorm 或 VSCode 打开 web 目录，在 Terminal(终端) 中，执行如下命令。
```
$ cd web
$ npm install
$ npm run serve
```

成功启动后，即可通过浏览器访问：http://localhost:8080/#/login

账号: admin 密码: 12345

**部署三：微信小程序（app）**

需要使用微信开发者工具打开 app 目录，在 Terminal(终端) 中，执行如下命令。
```
$ cd app 
$ npm install
```

在编译运行微信小程序之前，你需要进行以下设置：

在微信开发者工具右上角->【详情】->【本地设置】-> 选择【使用npm模块】和【不校验合法域名，web-view（业务域名）、TLS版本...】。最后，在微信开发者工具的工具栏->【工具】->【构建npm】。

**运行结果**

商家后台：

| | | |
|---|---|---|
| ![](https://github.com/zchengo/imall/blob/main/demo/res/w1.png) | ![](https://github.com/zchengo/imall/blob/main/demo/res/w2.png) | ![](https://github.com/zchengo/imall/blob/main/demo/res/w3.png) |
| ![](https://github.com/zchengo/imall/blob/main/demo/res/w4.png) | ![](https://github.com/zchengo/imall/blob/main/demo/res/w5.png) | ![](https://github.com/zchengo/imall/blob/main/demo/res/w6.png) |

小程序商城：

| | | | | | |
|---|---|---|---|---|---|
| ![](https://github.com/zchengo/imall/blob/main/demo/res/a1.png) | ![](https://github.com/zchengo/imall/blob/main/demo/res/a2.png) | ![](https://github.com/zchengo/imall/blob/main/demo/res/a3.png) | ![](https://github.com/zchengo/imall/blob/main/demo/res/a4.png) | ![](https://github.com/zchengo/imall/blob/main/demo/res/a5.png) | ![](https://github.com/zchengo/imall/blob/main/demo/res/a6.png) |


