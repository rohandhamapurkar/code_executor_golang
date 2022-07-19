package httpmiddlewares

import (
	"log"
	"net"
	"net/http"
	"rohandhamapurkar/code-executor/core/constants"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type visitor struct {
	limiter   *rate.Limiter
	createdAt time.Time
}

// Create a map to hold the rate limiters for each visitor and a mutex.
var visitors = make(map[string]*visitor)
var mu sync.Mutex

func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for ip, v := range visitors {
			// clear visitor after 1 minute from creation to restore burst
			if time.Since(v.createdAt) > time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

// Retrieve and return the rate limiter for the current visitor if it
// already exists. Otherwise create a new rate limiter and add it to
// the visitors map, using the IP address as the key.
func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(rate.Every(time.Minute), 15)
		// Include the current time when creating a new visitor.
		visitors[ip] = &visitor{limiter, time.Now()}
		return limiter
	}

	return v.limiter
}

func RateLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ip, _, err := net.SplitHostPort(ctx.Request.RemoteAddr)
		if err != nil {
			log.Println(err.Error())
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": constants.INTERNAL_SERVER_ERROR})
			return
		}

		limiter := getVisitor(ip)
		if limiter.Allow() == false {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": constants.USER_RATE_LIMITED})
			return
		}
		ctx.Next()
	}
}
