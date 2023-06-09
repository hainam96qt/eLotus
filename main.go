package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	configs "elotus/config"
	authentication2 "elotus/internal/endpoint/authentication"
	upload_file "elotus/internal/endpoint/file"
	"elotus/internal/service/authentication"
	upload_file2 "elotus/internal/service/file"
	"elotus/internal/service/jwt"
	"elotus/pkg/db/mysql_db"
	"elotus/pkg/midleware"
)

func main() {
	startedAt := time.Now()
	defer func() {
		log.Printf("application stopped after %s\n", time.Since(startedAt))
	}()

	conf, err := configs.NewConfig()
	if err != nil {
		log.Print(err)
	}

	globalCtx, glbCtxCancel := context.WithCancel(context.Background())

	httpSrv, err := initHTTPServer(globalCtx, conf)
	if err != nil {
		log.Panicf("failed to init http server %s \n", err)
	}

	go func() {
		log.Printf("starting HTTP server at: %s\n", conf.Server.Address)
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("failed to start HTTP server: %s \n", err)
		}
	}()

	// Keep the application running until signals trapped
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("%s signal trapped. Stopping application", <-sigChan)

	glbCtxCancel()
	// First terminate the HTTP gateway
	shutdownCtx, shutdownCtxCancel := context.WithTimeout(context.Background(), conf.Server.ShutdownTimeout)
	defer shutdownCtxCancel()
	if err := httpSrv.Shutdown(shutdownCtx); err != nil {
		log.Printf("failed to gracefully shutdown the HTTP gateway server: %s\n", err)
	} else {
		log.Println("HTTP gateway server stopped gracefully")
	}
}

func initHTTPServer(ctx context.Context, conf *configs.Config) (httpServer *http.Server, err error) {
	r := chi.NewRouter()
	midleware.AuthenticateMW = midleware.NewAuthenticateMiddleware(jwt.Algorithm.Alg(), conf.Jwt.SecretKey)

	// create endpoint here
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	dbConn, err := mysql_db.ConnectDatabase(conf.Mysqldb)
	if err != nil {
		log.Panicf("failed to connect database:: %s \n", err)
		return
	}

	// service
	jwtService := jwt.NewJwtService(conf)
	authSvc := authentication.NewAuthenticationService(conf, dbConn, jwtService)
	uploadFileSvc := upload_file2.NewUploadFileService(conf, dbConn)

	// handler
	authentication2.InitAuthenticationHandler(r, authSvc)
	upload_file.InitAuthenticationHandler(r, uploadFileSvc)

	return &http.Server{
		Addr:         conf.Server.Address,
		ReadTimeout:  conf.Server.ReadTimeout,
		WriteTimeout: conf.Server.WriteTimeout,
		Handler:      r,
	}, nil
}
