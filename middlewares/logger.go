package middlewares

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	logFilePath = ""
)

func init() {
	if dir, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		logFilePath = path.Join(dir, "logs")
	}

	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		panic(err)
	}
}

func Logger() *logrus.Logger {
	now := time.Now()

	logFileName := now.Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)

	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.999999999Z07:00",
	})

	return logger
}

func LogMessage(c *gin.Context, message string) {
	Logger().WithFields(logrus.Fields{
		"QueryID":     c.GetInt64(ContextKeyTraceID),
		"QueryUserID": c.GetUint(ContextKeyUserID),
		"QueryIP":     c.ClientIP(),
		"Path":        c.Request.RequestURI,
	}).Info(message)
}

func LogError(c *gin.Context, err error) {
	Logger().WithFields(logrus.Fields{
		"QueryID":     c.GetInt64(ContextKeyTraceID),
		"QueryUserID": c.GetUint(ContextKeyUserID),
		"QueryIP":     c.ClientIP(),
		"Path":        c.Request.RequestURI,
	}).Error(err.Error())
}

func LoggingToFiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		finish := time.Now()

		latency := finish.Sub(start)
		traceID := c.GetInt64(ContextKeyTraceID)
		method := c.Request.Method
		uri := c.Request.RequestURI
		status := c.Writer.Status()
		client := c.ClientIP()
		id := c.GetUint(ContextKeyUserID)

		Logger().Infof("| %v | %3d | %13v | %13v | %15s | %s | %s |",
			id,
			status,
			traceID,
			latency,
			client,
			method,
			uri,
		)
	}
}
