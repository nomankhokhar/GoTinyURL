package routes

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nomankhokhar/GoTonyUrl/api/database"
	"github.com/nomankhokhar/GoTonyUrl/api/models"
	"github.com/redis/go-redis/v9"
)

func ShortenURL(c *gin.Context) {
	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Parse JSON"})
	}

	r2 := database.CreateClient(1)
	defer r2.Close()

	_, err := r2.Get(database.Ctx, c.ClientIP()).Result()

	if err == redis.Nil {
		_ = r2.Set(database.Ctx, c.ClientIP(), os.Getenv("API_QUOTA"), 30*60*time.Second)
	} else {
		val, _ := r2.Get(database.Ctx, c.ClientIP()).Result()
		valInt, _ := strconv.Atoi(val)

		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.ClientIP()).Result()
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error":            "rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
			return
		}
	}

}
