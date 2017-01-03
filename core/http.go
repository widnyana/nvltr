package core

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/widnyana/nvltr/conf"
)

// Server hold the engine and tls cert.
type Server struct {
	Addr    string
	Engine  *gin.Engine
	TLSCert string
	TLSKey  string
}

func abortWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
	c.Abort()
}

func rootHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello!")
	c.Next()
}

func routerEngine() *gin.Engine {
	// set server mode
	gin.SetMode(conf.Config.Core.Mode)

	r := gin.Default()

	// Global middleware
	r.Use(LogMiddleware())
	r.GET("/", rootHandler)

	// register handler
	for _, endp := range routeProvider() {
		r.Handle(endp.method, endp.path, endp.handler)
	}

	return r
}

// MakeHTTPServer create instance of HTTP Server
func MakeHTTPServer() *Server {
	s := &Server{
		Addr:   fmt.Sprintf("%s:%s", conf.Config.Core.Addr, conf.Config.Core.Port),
		Engine: routerEngine(),
	}

	if conf.Config.Core.SSL {
		s.TLSKey = TLSKey
		s.TLSCert = TLSCert
	}

	return s
}

// RunServer will run with default gin server
func (s *Server) RunServer() error {
	var err error

	if conf.Config.Core.SSL {
		err = s.Engine.RunTLS(s.Addr, conf.Config.Core.CertPath, conf.Config.Core.KeyPath)
	} else {
		err = s.Engine.Run(s.Addr)
	}

	return err
}

// RunEndlessHTTPServer provide run http or https protocol.
func (s *Server) RunEndlessHTTPServer() error {
	var err error
	if conf.Config.Core.SSL {
		server := NewServer(s.Addr, s.Engine)
		err = server.ListenAndServeTLSfromString(s.TLSCert, s.TLSKey)

	} else {
		err = ListenAndServe(":"+conf.Config.Core.Port, routerEngine())
	}

	err = s.Engine.Run("localhost:8088")
	return err
}

// func (s *Server) RunHTTPServer() error {
// 	var err error
// 	if Config.Core.SSL {
// 		LogAccess.Info("Listening on TLS")
// 		server := NewServer(s.Addr, s.Engine)
// 		err = server.ListenAndServeTLSfromString(s.TLSCert, s.TLSKey)
// 	} else {
// 		LogAccess.Info("Listening on HTTP")
// 		err = ListenAndServe(":"+Config.Core.Port, routerEngine())
// 	}

// 	return err
// }
