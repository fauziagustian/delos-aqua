package main

import (
	"fmt"
	"os"

	config "github.com/fauziagustian/delos-aqua/config"
	"github.com/fauziagustian/delos-aqua/internal/handler"
	"github.com/fauziagustian/delos-aqua/internal/repository"
	"github.com/fauziagustian/delos-aqua/internal/route"
	"github.com/fauziagustian/delos-aqua/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDb()

	//connection db to repository
	userRepository := repository.NewUserRepository(db)
	farmRepository := repository.NewFarmRepository(db)
	pondRepository := repository.NewPondRepository(db)
	userAgentRepository := repository.NewUserAgentRepository(db)

	//connection service to handler
	userService := service.NewUserService(&service.USConfig{UserRepository: userRepository})
	authService := service.NewAuthService(&service.ASConfig{UserRepository: userRepository})
	jwtService := service.NewJWTService(&service.JWTSConfig{})
	farmService := service.NewFarmService(&service.FConfig{FarmRepository: farmRepository})
	pondService := service.NewPondService(&service.PConfig{PondRepository: pondRepository})
	userAgentService := service.NewUserAgentService(&service.UAConfig{UserAgentRepository: userAgentRepository})

	//connection handler to package route
	h := handler.NewHandler(&handler.HandlerConfig{
		UserService:      userService,
		AuthService:      authService,
		JWTService:       jwtService,
		FarmService:      farmService,
		PondService:      pondService,
		UserAgentService: userAgentService,
	})
	routes := route.NewRouter(&route.RouterConfig{UserService: userService, JWTService: jwtService, FarmService: farmService, PondService: pondService})

	router := gin.Default()
	//i'm not using swagger for api docs, so i coment this
	//router.Static("/docs", "./pkg/swaggerui")
	router.NoRoute(h.NoRoute)

	version := os.Getenv("API_VERSION")
	api := router.Group(fmt.Sprintf("/v1/%s", version))

	//this is set up for the route
	routes.Auth(api, h)
	routes.User(api, h)
	routes.Farm(api, h)
	routes.Pond(api, h)
	routes.UserAgent(api, h)

	router.Run(":8080")
}
