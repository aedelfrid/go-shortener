package main

import (
	"context"
	"net/http"
	"net/url"
	"sync/atomic"

	"fmt"
	"math/rand"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var ctx = context.Background()
var rdb *redis.Client

var counter int64 = 10000

func randomID() string {
	return fmt.Sprintf("%d", rand.Intn(999999))
}

func init() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, using default system envs")
	}

	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))

	rdb = redis.NewClient(opt)
}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
	})

	r.POST("/shorten", func(c *gin.Context) {
		var input struct {
			LongURL string `json:"url"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		parsedURL, err := url.ParseRequestURI(input.LongURL)

		if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid URL. Please include http:// or https://",
			})
			return
		}

		newID := atomic.AddInt64(&counter, 1)

		shortCode := Encode(uint64(newID))

		err = rdb.Set(ctx, shortCode, input.LongURL, 0).Err()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save to Redis"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"short_url": "http://localhost:4000/" + shortCode,
			"code":      shortCode,
		})
	})

	r.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")

		longURL, err := rdb.Get(ctx, code).Result()

		if err == redis.Nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Short link not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, longURL)

	})

	r.Run(os.Getenv("PORT"))
}
