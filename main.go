package main

import (
	"fmt"
	"net/http"
	"time"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/etashkinov/hall-of-fame/middlewares"
	"github.com/etashkinov/hall-of-fame/routes"
)

func main() {
	router := gin.Default()

	initialGinConfig(router)
	router.Use(middlewares.GinBodyLogMiddleware)
	routes.ApplicationV1Router(router)
	startServer(router)

}

func initialGinConfig(router *gin.Engine) {
	router.Use(limit.MaxAllowed(200))
	router.Use(cors.Default())
}

func startServer(router http.Handler) {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error in config file: %s \n", err))
	}
	serverPort := fmt.Sprintf(":%s", viper.GetString("ServerPort"))
	s := &http.Server{
		Addr:           serverPort,
		Handler:        router,
		ReadTimeout:    18000 * time.Second,
		WriteTimeout:   18000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(fmt.Errorf("fatal error description: %s \n", err))
	}
}
