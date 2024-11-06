package calculate

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

// Initialize Redis client
func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Assuming Redis is running locally or in Docker with the correct hostname
	})
}

// Check if sum result is cached in Redis
func GetCachedSum(a, b int) (int, bool) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("sum:%d:%d", a, b)

	// Try to get the cached value
	val, err := rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		// Cache miss
		return 0, false
	} else if err != nil {
		log.Fatalf("Error fetching from cache: %v", err)
		return 0, false
	}

	// Cache hit, return the cached sum
	cachedSum, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Error converting cached value to int: %v", err)
		return 0, false
	}

	return cachedSum, true
}

// Store sum result in Redis cache
func SetCacheSum(a, b, sum int) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("sum:%d:%d", a, b)
	// Set the sum in the cache with an expiration time of 1 hour
	err := rdb.Set(ctx, cacheKey, sum, time.Hour).Err()
	if err != nil {
		log.Fatalf("Error setting cache: %v", err)
	}
}

// Sum function that uses cache
func Sum(a, b int) int {
	// Check if the sum is already cached
	cachedSum, found := GetCachedSum(a, b)
	if found {
		// Return the cached sum
		log.Println("Using cached result")
		return cachedSum
	}

	// If not cached, calculate the sum and cache it
	sum := a + b
	SetCacheSum(a, b, sum)
	log.Println("Calculating sum and caching the result")
	return sum
}
