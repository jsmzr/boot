# boot
[![Build Status](https://github.com/jsmzr/boot/workflows/Run%20Tests/badge.svg?branch=main)](https://github.com/jsmzr/boot/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/jsmzr/boot/branch/main/graph/badge.svg?token=HNQCAN3UVR)](https://codecov.io/gh/jsmzr/boot)

Integrate commonly used libraries in Golang, reduce the threshold of dependency management, and improve the efficiency of project construction and docking of various libraries.

Based on the principle that convention is greater than configuration, reduce configuration management work. Combined with anonymous import to complete the registration and initialization of the library, focus more on business.

## Getting started

Taking the GIN Web project as an example, the sequence in which a project is created based on ***boot*** is as follows

1. Create a directory `mkdir boot-demo && cd boot-demo`, initialize `go mod init github.com/jsmzr/boot-demo`
2. Pull dependencies, `go get -u github.com/jsmzr/boot`
3. Import the plug-ins you need to use in `main.go`
4. update dependencies `go mod tidy`
5. Complete the interface logic writing, add swagger doc
6. Write routing rules
7. register router by `func init` in `router.go`
8. run `swag init` generate doc, and import `docs` in `main.go`
9. use `bootGin.Run()` start

```golang
// main.go
package main

import (
	"fmt"

	bootGin "github.com/jsmzr/boot/gin"
    // swagger plugin 
	_ "github.com/jsmzr/boot/gin/swagger"
    // logrus plugin
	_ "github.com/jsmzr/boot/log/logrus"

    // use swag init generate
	_ "github.com/jsmzr/boot-demo/docs"
    // gin router register
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