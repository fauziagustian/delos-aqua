package route

import (
	"github.com/fauziagustian/delos-aqua/internal/handler"
	"github.com/fauziagustian/delos-aqua/internal/middleware"
	s "github.com/fauziagustian/delos-aqua/internal/service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	userService s.UserService
	jwtService  s.JWTService
	farmService s.FarmService
}

type RouterConfig struct {
	UserService s.UserService
	JWTService  s.JWTService
	FarmService s.FarmService
}

func NewRouter(c *RouterConfig) *Router {
	return &Router{
		userService: c.UserService,
		jwtService:  c.JWTService,
		farmService: c.FarmService,
	}
}

func (r *Router) Auth(route *gin.RouterGroup, h *handler.Handler) {
	route.POST("/sign-up", h.Register)
	route.POST("/sign-in", h.Login)
}

func (r *Router) User(route *gin.RouterGroup, h *handler.Handler) {
	route.GET("/profiles", middleware.AuthMiddleware(r.jwtService, r.userService), h.Profile)
}

func (r *Router) Farm(route *gin.RouterGroup, h *handler.Handler) {
	route.GET("/farm", middleware.AuthMiddleware(r.jwtService, r.userService), h.GetFarm)
	route.POST("/farm", middleware.AuthMiddleware(r.jwtService, r.userService), h.CreateFarm)
	route.PUT("/farm/:id", middleware.AuthMiddleware(r.jwtService, r.userService), h.UpdateFarm)
	route.DELETE("/farm/:id", middleware.AuthMiddleware(r.jwtService, r.userService), h.DeleteFarm)
}
