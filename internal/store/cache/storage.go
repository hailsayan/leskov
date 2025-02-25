package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/hailsayan/leskov/internal/types"
)

type IUsers interface {
	Get(context.Context, int) (*types.User, error)
	Set(context.Context, *types.User) error
}

type IProducts interface {
	Get(context.Context, int) (*types.Product, error)
	Set(context.Context, *types.Product) error
	GetAll(ctx context.Context) ([]*types.Product, error)
	SetAll(ctx context.Context, products []*types.Product) error
}

type Storage struct {
	Users    IUsers
	Products IProducts
}

func NewRedisStorage(rdb *redis.Client) Storage {
	return Storage{
		Users:    &UserStore{rdb: rdb},
		Products: &ProductStore{rdb: rdb},
	}
}
