package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --------------------------------------------------------------------------------------
func (h *Handler) sendSms(c *gin.Context) {

	sendSmsService := (*h.environment).SendSms
	receivedPhone := c.Param("phone")

	result, err := (*sendSmsService).Send(fmt.Sprint(receivedPhone), "ccc")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, gin.MIMEJSON, result)
}
