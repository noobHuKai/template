package admin_controller

import (
	"github.com/gin-gonic/gin"
	response "github.com/noobHuKai/internal/models/response_model"
)

// Login godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Login(c *gin.Context) {
	response.Ok(c)
}
