package controller

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

var (
	serverAddr = ""
	serverPort = ""
)

func ProxyGenerator() gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse("http://" + serverAddr + ":" + serverPort)
		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
