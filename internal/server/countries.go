package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// --------------------------------------------------------------------------------------
func (h *Handler) getCountries(c *gin.Context) {
	// request countries from cache
	data, err := (*h.environment.Countries).GetCountries()
	if err != nil {
		// countries not found
		(*h.environment.Logger).Warningln("List of countries not found in cache", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, gin.MIMEJSON, data)
}
