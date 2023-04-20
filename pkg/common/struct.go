package common

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Global struct {
	DbMyTodo *gorm.DB
	Redis    *redis.Client
}
