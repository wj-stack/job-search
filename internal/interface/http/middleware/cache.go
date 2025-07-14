package middleware

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type Cache struct {
	store map[string]string
	mu    sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]string),
	}
}

// CacheMiddleware 缓存中间件示例
func CacheMiddleware(cache *Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.URL.String()
		cache.mu.Lock()
		if val, exists := cache.store[key]; exists {
			c.String(200, val)
			c.Abort()
			cache.mu.Unlock()
			return
		}
		cache.mu.Unlock()

		c.Next()

		if c.Writer.Status() == 200 {
			cache.mu.Lock()
			cache.store[key] = c.GetString("response")
			cache.mu.Unlock()
		}
	}
}
