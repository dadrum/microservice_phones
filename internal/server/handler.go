package server

import (
	"micro_service_phone/internal/environment"

	"github.com/gin-gonic/gin"
)

type Countries interface {
	GetCountries() (bool, error)
}

type Handler struct {
	environment *environment.Environment
	Countries
}

// --------------------------------------------------------------------------------------
// handler consructor with dependencies
func NewHandler(environment *environment.Environment) *Handler {
	return &Handler{environment: environment}
}

// --------------------------------------------------------------------------------------
// routes initialization in handler
func (h *Handler) InitRoutes() *gin.Engine {
	defer h.environment.Logger.Debugln("server routes initialization")
	router := gin.New()

	v1 := router.Group("/v1")
	{
		v1.GET("/heardbeat", h.heardbeat)

		v1.GET("/countries", h.getCountries)

		v1.GET("/validate/:phone", h.validatePhone)

		v1.POST("/send-message/:phone", h.sendSms)
	}

	return router
}
