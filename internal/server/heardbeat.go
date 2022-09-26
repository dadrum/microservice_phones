package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// --------------------------------------------------------------------------------------
func (h *Handler) heardbeat(c *gin.Context) {
	c.String(http.StatusNoContent, "")
}
