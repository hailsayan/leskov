package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hailsayan/woland/internal/types"
)

type ProductStore struct {
	rdb *redis.Client
}

const ExpTime = time.Minute

func (p *ProductStore) Get(ctx context.Context, id int) (*types.Product, error) {
	key := fmt.Sprintf("product:%d", id)
	data, err := p.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var product types.Product
	if err := json.Unmarshal([]byte(data), &product); err != nil {
		return nil, err
	}

	fmt.Println("Cache Hit: Product retrieved from cache")
	return &product, nil
}

func (p *ProductStore) Set(ctx context.Context, product *types.Product) error {
	key := fmt.Sprintf("product:%d", product.ID)
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}

	return p.rdb.Set(ctx, key, data, ExpTime).Err()
}
