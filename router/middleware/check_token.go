package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"opsHeart_agent/common"
	"opsHeart_agent/conf"
)

func TokenChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check uuid and token
		reqUUID, token, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "auth required",
			})
			c.Abort()
			return
		}

		if reqUUID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "uuid blank",
			})
			c.Abort()
			return
		}

		right := false
		for _, v := range conf.GetSvrUUIDs() {
			if v == reqUUID {
				right = true
				break
			}
		}
		if !right {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "uuid wrong",
			})
			c.Abort()
			return
		}

		if token != common.SelfToken {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token wrong",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
