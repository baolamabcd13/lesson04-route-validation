package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	limiter  *rate.Limiter
	lastseen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*Client)
)

func getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}
	return ip
}

func getRateLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	client, exists := clients[ip]
	if !exists {
		limiter := rate.NewLimiter(5, 15) // 5 request/sec, brust 10: số lượng request tối đa có thể xử lí
		newClient := &Client{limiter, time.Now()}
		clients[ip] = newClient

		return limiter
	}

	client.lastseen = time.Now()
	return client.limiter
}

func CleanUpClients() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()

		for ip, client := range clients {
			if time.Since(client.lastseen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

// ab -n 20 -c 1 -H "X-API-Key:20019266-db0b-4036-a694-89f2ca9e3e8f" http://localhost:8080/api/v1/categories/golang
func RateLimitingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getClientIP(ctx)
		//log.Println("IP Address:", ip) // ::1 <=> 127.0.0.1
		limiter := getRateLimiter(ip)
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "Too many request",
				"message": "Bạn đã gửi quá nhiều request, hãy thử lại sau!!! ",
			})
			return
		}
		ctx.Next()
	}
}
