package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type client_usr struct {
	request   int
	last_seen time.Time
}

var (
	clients   = make(map[string]*client_usr)
	mu        sync.Mutex
	rateLimit = 5
	window    = 1 * time.Minute
)

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		client_ip := c.ClientIP()
		mu.Lock()
		client, ok := clients[client_ip]
		if !ok {
			client = &client_usr{request: 1, last_seen: time.Now()}
			clients[client_ip] = client
		} else {
			if time.Since(client.last_seen) > window {
				client.request = 1
				client.last_seen = time.Now()
			} else {
				client.request++
			}
		}
		mu.Unlock()
		if client.request > rateLimit {
			c.JSON(429, gin.H{"msg": "Too many request try after sometime"})
			c.Abort()
			return
		}
		c.Next()
	}
}
