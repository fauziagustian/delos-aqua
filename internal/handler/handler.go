package handler

import (
	"net/http"

	s "github.com/fauziagustian/delos-aqua/internal/service"
	"github.com/fauziagustian/delos-aqua/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService      s.UserService
	authService      s.AuthService
	jwtService       s.JWTService
	farmService      s.FarmService
	pondService      s.PondService
	userAgentService s.UserAgentService
}

type HandlerConfig struct {
	UserService      s.UserService
	AuthService      s.AuthService
	JWTService       s.JWTService
	FarmService      s.FarmService
	PondService      s.PondService
	UserAgentService s.UserAgentService
}

func NewHandler(c *HandlerConfig) *Handler {
	return &Handler{
		userService:      c.UserService,
		authService:      c.AuthService,
		jwtService:       c.JWTService,
		farmService:      c.FarmService,
		pondService:      c.PondService,
		userAgentService: c.UserAgentService,
	}
}

func (h *Handler) NoRoute(c *gin.Context) {
	response := utils.ErrorResponse("page not found", http.StatusNotFound, "page not found")
	c.JSON(http.StatusNotFound, response)
}
