package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/hailsayan/woland/internal/types"
)

type IUsers interface {
	Get(context.Context, int) (*types.User, error)
	Set(context.Context, *types.User) error
}

type Storage struct {
	Users IUsers
}

func NewRedisStorage(rdb *redis.Client) Storage {
	return Storage{
		Users: &UserStore{rdb: rdb},
	}
}
