package calculate

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis(redisAddr string) {
	rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr, // e.g., "redis:6379"
	})
}

func Sum(a, b int) int {
	// Redis key
	cacheKey := fmt.Sprintf("sum:%d:%d", a, b)

	// Try to fetch the result from Redis
	cachedResult, err := rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		// Cache hit: Return cached result
		cachedSum, _ := strconv.Atoi(cachedResult)
		return cachedSum
	}

	// Cache miss: Calculate the sum
	sum := a + b

	// Store the result in Redis
	err = rdb.Set(ctx, cacheKey, fmt.Sprintf("%d", sum), 0).Err()
	if err != nil {
		fmt.Println("Error setting cache:", err)
	}

	return sum
}
