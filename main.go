package main

import (
	"book-store/database"
	"book-store/datastores"
	"book-store/handlers"
	middleware "book-store/middlewares"
	"book-store/services"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// application context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	// mysql connection initialization
	database.ConnectMysqlDB(ctx)

	// dependency injection
	store := datastores.New(database.GetConnection())
	svc := services.New(store)
	controller := handlers.New(svc)

	app := gin.New()

	apiGrp := app.Group("/v1")

	apiGrp.Use(middleware.ExceptionHandler())
	//Routes
	apiGrp.POST("/book", controller.Create)
	apiGrp.GET("/book/:id", controller.FindById)
	apiGrp.GET("/book", controller.FindAll)
	apiGrp.PATCH("book/:id", controller.Update)
	apiGrp.DELETE("book/:id", controller.Delete)

	server := &http.Server{
		Addr:    ":" + viper.GetString("HTTP_PORT"),
		Handler: app,
	}
	// interrupts
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		log.Println("app to be stopped")

		cancel()
		database.Close()
		server.Shutdown(context.Background())
	}()

	err := server.ListenAndServe()
	if err != nil {
		cancel()
	}

}
