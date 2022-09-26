package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --------------------------------------------------------------------------------------
func (h *Handler) validatePhone(c *gin.Context) {
	phoneValidator := (*h.environment).Phonevalidator
	receivedPhone := c.Param("phone")

	result, err := (*phoneValidator).Validate(fmt.Sprint(receivedPhone))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, gin.MIMEJSON, result)
}
