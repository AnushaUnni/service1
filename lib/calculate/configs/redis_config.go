package configs

type RedisConfig struct {
	RedisAddr string // Redis server address
}

func NewRedisConfig(redisAddr string) *RedisConfig {
	return &RedisConfig{
		RedisAddr: redisAddr,
	}
}
