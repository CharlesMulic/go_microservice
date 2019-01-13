package main

import (
	"log"
	"net/http"
	"os"

	"github.com/CharlesMulic/go_microservice/homepage"
	"github.com/CharlesMulic/go_microservice/server"
)

var (
	CertFile    = os.Getenv("GO_CERT_FILE")
	KeyFile     = os.Getenv("GO_KEY_FILE")
	ServiceAddr = os.Getenv("GO_SERVICE_ADDR")
)

func main() {
	logger := log.New(os.Stdout, "Go Microservice ", log.LstdFlags|log.Lshortfile)
	mux := http.NewServeMux()
	h := homepage.NewHandlers(logger)

	mux.HandleFunc("/", h.Home)

	// srv := NewServer(mux, ServiceAddr)
	srv := server.New(mux, ":8080")

	logger.Println("Server Starting...")
	// err := http.ListenAndServe(":8080", mux)
	// err := srv.ListenAndServeTLS(CertFile, KeyFile)
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
