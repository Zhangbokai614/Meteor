package main

import (
	"github.com/Zhangbokai614/go-template/middlewares"
	"github.com/Zhangbokai614/go-template/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.AttachTraceID())
	r.Use(middlewares.LoggingToFiles())

	r.Use(cors.Default())

	routers.Init(r)

	r.Run(":6145") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
