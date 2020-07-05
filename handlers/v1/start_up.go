package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"opsHeart_agent/common"
	"opsHeart_agent/logger"
	"opsHeart_agent/service/start_up"
)

func HandleStartUp(c *gin.Context) {
	// bind data
	var s start_up.RegToken
	err := c.ShouldBindJSON(&s)
	if err != nil {
		logger.AgentLog.Errorf("action=server register feedback;do=bind data;err=%s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	// handle status
	err = s.HandleStatus()
	if err != nil {
		logger.AgentLog.Errorf("action=server register feedback;do=handle status;err=%s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})

	// stop register router
	common.StopRegister <- struct{}{}
}
