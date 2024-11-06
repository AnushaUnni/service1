package configs

type RedisConfig struct {
	RedisAddr     string // Redis server address
	RedisPassword string // Redis password (if needed)
	RedisDB       int    // Redis database index (default is 0)
	PoolSize      int    // Number of connections in the Redis connection pool
}

func NewRedisConfig(redisAddr string, redisPassword string, redisDB int, poolSize int) *RedisConfig {
	return &RedisConfig{
		RedisAddr:     redisAddr,
		RedisPassword: redisPassword,
		RedisDB:       redisDB,
		PoolSize:      poolSize,
	}
}
