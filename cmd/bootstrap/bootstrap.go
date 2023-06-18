package bootstrap

import (
	"context"
	"log"
	"net/http"

	auth_handler "github.com/krlspj/mind-sprint-be/internal/auth/platform/http"
	"github.com/krlspj/mind-sprint-be/internal/auth/platform/storage/inmemory"
	auth_service "github.com/krlspj/mind-sprint-be/internal/auth/service"
	"github.com/krlspj/mind-sprint-be/internal/server"
)

const (
	// Database
	dbType     = "inmemory" // "mongo" | "inmemory"
	dbName     = "mind_sprint"
	userColl   = "users"
	configColl = "config"
	port       = "6000"
)

// Run starts the service.
// Does the initialization of different packages and inject to corresponding services
func Run() error {

	authUserRepo := inmemory.NewUserRepositoryStub()
	authService := auth_service.NewAuthService(authUserRepo)
	authHandler := auth_handler.NewAuthHandler(authService)

	//	mux := http.NewServeMux()
	//	mux.HandleFunc("/health", healthHandler)
	//
	//	return http.ListenAndServe(httpAddr, mux)

	srv := server.NewServer(context.TODO(), "", port, authHandler)

	return srv.RunHTTPServer()

}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("healthcheck log:", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

}
