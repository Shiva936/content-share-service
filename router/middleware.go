package router

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("X-Request-Id", uuid.New().String())

		t := time.Now()

		c.Next()

		latency := time.Since(t)

		log.Printf("%s : %s : %s : %s : %d : %s : %s\n",
			"INFO",
			c.Request.Header.Get("X-Request-Id"),
			c.Request.Method,
			c.Request.RequestURI,
			c.Writer.Status(),
			latency,
			c.Request.Proto,
		)
	}
}
