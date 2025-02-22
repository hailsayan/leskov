package api

import (
	"context"
	"fmt"

	configs "github.com/hailsayan/woland/internal/config"
	"github.com/hailsayan/woland/internal/types"
)

func (s *Server) getProductsFromCacheOrDB(ctx context.Context, id *int) (interface{}, error) {
	if !configs.Envs.Enabled {
		if id != nil {
			return s.store.Product.GetProductByID(*id)
		}
		return s.store.Product.GetProducts()
	}

	if id != nil {
		product, err := s.cacheStorage.Products.Get(ctx, *id)
		if err != nil {
			return nil, err
		}

		if product == nil {
			product, err = s.store.Product.GetProductByID(*id)
			if err != nil {
				return nil, err
			}

			if err := s.cacheStorage.Products.Set(ctx, product); err != nil {
				fmt.Println("error saving the product in cache:", err)
			}
		}
		return product, nil
	}
	products, err := s.cacheStorage.Products.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if products == nil {
		products, err = s.store.Product.GetProducts()
		if err != nil {
			return nil, err
		}

		if err := s.cacheStorage.Products.SetAll(ctx, products); err != nil {
			fmt.Println("error saving the products list in cache:", err)
		}
	}

	return products, nil
}

func (s *Server) getUser(ctx context.Context, id int) (*types.User, error) {
	if !configs.Envs.Enabled {
		return s.store.Users.GetUserByID(id)
	}

	user, err := s.cacheStorage.Users.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		user, err = s.store.Users.GetUserByID(id)
		if err != nil {
			return nil, err
		}

		if err := s.cacheStorage.Users.Set(ctx, user); err != nil {
			return nil, err
		}
	}

	return user, nil
}
