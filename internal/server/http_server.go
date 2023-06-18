package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	auth_handler "github.com/krlspj/mind-sprint-be/internal/auth/platform/http"
)

const (
	SERVER_GIN = "gin"
	SERVER_GO  = "native"
)

type httpServer struct {
	httpAddr string
	engine   *gin.Engine

	//mux      http.Handler

	// Handlers
	authHandler auth_handler.AuthHandler
	//	az *authz_handler.AuthzMidd
	//ms       mid.MiddlewareService

	//mux      *pat.PatternServeMux
	//mux      *http.ServeMux
}

func NewServer(
	ctx context.Context,
	host, port string,
	ah auth_handler.AuthHandler,
) httpServer {

	svr := httpServer{
		httpAddr:    fmt.Sprintf(host + ":" + port),
		engine:      gin.Default(),
		authHandler: ah,
	}
	return svr
}

func (s *httpServer) GetEngine() *gin.Engine {
	return s.engine
}

func (s httpServer) RunHTTPServer() error {

	s.registerRoutes()

	log.Printf("Listening on %s\n", s.httpAddr)

	return s.engine.Run(s.httpAddr)

}

func (s httpServer) registerRoutes() {
	log.Println("engine routes...")

	s.engine.GET("/health", healthCheck())

	// Auth
	s.engine.POST("/register", s.authHandler.RegisterNewUser)
	s.engine.POST("/login", s.authHandler.Login)

}

func healthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Server is up and running!")
	}
}
