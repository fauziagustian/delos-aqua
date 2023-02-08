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
	pondService s.PondService
}

type RouterConfig struct {
	UserService s.UserService
	JWTService  s.JWTService
	FarmService s.FarmService
	PondService s.PondService
}

func NewRouter(c *RouterConfig) *Router {
	return &Router{
		userService: c.UserService,
		jwtService:  c.JWTService,
		farmService: c.FarmService,
		pondService: c.PondService,
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

func (r *Router) Pond(route *gin.RouterGroup, h *handler.Handler) {
	route.GET("/pond", middleware.AuthMiddleware(r.jwtService, r.userService), h.GetPond)
	route.POST("/pond", middleware.AuthMiddleware(r.jwtService, r.userService), h.CreatePond)
	route.PUT("/pond/:id", middleware.AuthMiddleware(r.jwtService, r.userService), h.UpdatePond)
	route.DELETE("/pond/:id", middleware.AuthMiddleware(r.jwtService, r.userService), h.DeletePond)
}
