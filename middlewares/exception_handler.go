package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {

				resp := ErrorResp{Error: r.(error).Error(), Desc: "Unknown Error"}

				log.Println("Panic recovered: ", resp)
				c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
			}
		}()
		c.Next()
	}
}

type ErrorResp struct {
	Error interface{} `json:"error,omitempty"`
	Desc  string      `json:"desc,omitempty"`
}
