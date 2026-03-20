package main

import (
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	//"github.com/go-redis/redis/v8"
	"fmt"
	"math/rand"
)

var urlStore = make(map[string]string)
var counter int64 = 10000

func randomID() string {
	return fmt.Sprintf("%d", rand.Intn(999999))
}

func main() {
	r := gin.Default()

	r.POST("/shorten", func(ctx *gin.Context) {
		var input struct {
			LongURL string `json:"url"`
		}

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		newID := atomic.AddInt64(&counter, 1)

		shortCode:= Encode(uint64(newID))

		urlStore[shortCode] = input.LongURL

		ctx.JSON(http.StatusOK, gin.H{
			"short_url": "http://localhost:4000/" + shortCode,
			"code" : shortCode,
		})
	})

	r.GET("/:code", func(ctx *gin.Context) {
		code := ctx.Param("code")
		LongURL, exists := urlStore[code]

		if !exists {
			ctx.JSON(http.StatusNotFound, gin.H{"error":"URL not found"})
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, LongURL)

	})

	r.Run(":4000")
}