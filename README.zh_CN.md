# boot
[![Build Status](https://github.com/jsmzr/boot/workflows/Run%20Tests/badge.svg?branch=main)](https://github.com/jsmzr/boot/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/jsmzr/boot/branch/main/graph/badge.svg?token=HNQCAN3UVR)](https://codecov.io/gh/jsmzr/boot)

整合 golang 种常用的库，降低依赖管理门槛，提高项目搭建及各类库对接效率。

基于**约定大于配置**的原则，减少配置管理工作。结合匿名导入完成库的注册、初始化，更专注于业务。


## 快速开始

以 gin web 项目为例，基于 boot 创建一个项目的顺序如下

1. 新建项目目录`mkdir boot-demo && cd boot-demo`，并完成初始化 `go mod init github.com/jsmzr/boot-demo`
2. 拉取 boot 依赖，`go get -u github.com/jsmzr/boot`
3. 在 `main.go` 中通过 import 声明需要用到的插件
4. 执行 `go mod tidy` 拉取依赖
5. 完成接口逻辑开发，并按照 swagger 规范编写接口文档
6. 在业务模块对应的路由文件中编写路由规则
7. 在全局路由文件 `router.go` 的 `init` 方法处把路由注册到 bootGin 中
8. 执行 `swag init` 生成接口文档，并在 `main.go` 中导入生成的 `docs` 
9. 通过 `bootGin.Run()` 启动后台服务


```golang
// main.go
package main

import (
	"fmt"

	bootGin "github.com/jsmzr/boot/gin"
    // swagger 插件声明
	_ "github.com/jsmzr/boot/gin/swagger"
    // logrus 插件声明
	_ "github.com/jsmzr/boot/log/logrus"

    // 执行 swag init 后生成的接口文档文件夹
	_ "github.com/jsmzr/boot-demo/docs"
    // gin 路由注册
	_ "github.com/jsmzr/boot-demo/router"
)

func main() {
	if err := bootGin.Run(); err != nil {
		fmt.Println(err)
	}
}

// router/router.go
package router

import bootGin "github.com/jsmzr/boot/gin"

func init() {
	bootGin.RegisterRouter(InitDemo)
}

// router/demo.go
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jsmzr/boot-demo/controller"
)

func InitDemo(e *gin.Engine) {
	e.GET("/hello", controller.Hello)
}

// controller/demo.go
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Tags demo
// @Summary hello world
// @Router /hello [get]
func Hello(c *gin.Context) {
	logrus.Info("request /hello")
	c.String(http.StatusOK, "hello world!")
}

```

当前支持的插件如下，boot 中已支持的插件从库的目录即可查看，也可以获取其他库的插件或自定义插件。

- [ ] log
    - [x] logrus
    - [ ] zap
- [ ] config
    - [x] apollo
- [ ] api document
    - [x] swagger(gin)
- [ ] trace
    - [x] skywalking(gin)
- [ ] mertrics
    - [x] prometheus
- [ ] database
    - [ ] mysql
    - [ ] oracle
- [ ] cache
    - [ ] redis

## 开发说明

首先基于 viper 管理所有配置，库初始化配置和业务配置均从 viper 获取即可，需要注意 viper 中定义的配置**优先级**。

其次以插件的方式进行各类库的注册、初始化管理。

### 插件说明

当前插件接口定义了三个方法，`Load`, `Enabled`, `Order`, 分别定义加载、开关、初始化顺序。

通常插件中会还需要一些默认配置，用于减少开发者的配置维护成本。