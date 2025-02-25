package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hailsayan/leskov/internal/types"
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

func (p *ProductStore) GetAll(ctx context.Context) ([]*types.Product, error) {
	keys, err := p.rdb.Keys(ctx, "product:*").Result()
	if err != nil {
		return nil, err
	}

	var products []*types.Product
	for _, key := range keys {
		data, err := p.rdb.Get(ctx, key).Result()
		if err == redis.Nil {
			continue
		} else if err != nil {
			return nil, err
		}

		var product types.Product
		if err := json.Unmarshal([]byte(data), &product); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (p *ProductStore) SetAll(ctx context.Context, products []*types.Product) error {
	for _, product := range products {
		key := fmt.Sprintf("product:%d", product.ID)
		data, err := json.Marshal(product)
		if err != nil {
			return err
		}

		if err := p.rdb.Set(ctx, key, data, ExpTime).Err(); err != nil {
			return err
		}
	}

	return nil
}
