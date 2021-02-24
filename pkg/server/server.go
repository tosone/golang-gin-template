package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tosone/logging"
)

// Initialize initialize
func Initialize() (err error) {
	if viper.GetBool("Debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	var app = gin.Default()

	app.Use(gin.Recovery())
	app.Use(gin.Logger())
	app.Use(cors.Default())

	app.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	var rateLimitMiddleware gin.HandlerFunc
	if rateLimitMiddleware, err = rateLimit(); err != nil {
		logging.Fatalf("create gin rate limit middleware with error: %v", err)
	}

	app.Use(rateLimitMiddleware)
	app.Use(gzip.Gzip(gzip.DefaultCompression))

	routers(app)

	app.NoRoute(func(context *gin.Context) {
		context.Status(http.StatusNotFound)
	})

	fmt.Printf("Listening and serving HTTP on http://127.0.0.1:%s\n", viper.GetString("ServerPort"))

	var srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("ServerPort")),
		Handler: app,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Fatalf("listen on %s with error: %v", viper.GetString("ServerPort"), err)
		}
	}()

	var quit = make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logging.Info("server is shutting down")

	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var serverShutdown = make(chan bool)
	go func() {
		if err = srv.Shutdown(ctx); err != nil {
			logging.Fatalf("server shutdown with error: %v", err)
		}
		serverShutdown <- true
	}()

	select {
	case <-ctx.Done():
		logging.Info("server cannot shutdown in 5 seconds")
	case <-serverShutdown:
		logging.Info("server shutdown graceful")
	}

	return
}
