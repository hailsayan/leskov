package api

import (
	"context"
	"fmt"

	configs "github.com/hailsayan/woland/internal/config"
	"github.com/hailsayan/woland/internal/types"
)

func (s *Server) getProduct(ctx context.Context, id int) (*types.Product, error) {
	if !configs.Envs.Enabled {
		return s.store.Product.GetProductByID(ctx, id)
	}

	product, err := s.cacheStorage.Products.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		product, err = s.store.Product.GetProductByID(ctx, id)
		if err != nil {
			return nil, err
		}

		if err := s.cacheStorage.Products.Set(ctx, product); err != nil {
			fmt.Println("error saving the product in cache:", err)
		}
	}

	return product, nil
}

func (s *Server) getUser(ctx context.Context, id int) (*types.User, error) {
	if !configs.Envs.Enabled {
		return s.store.Users.GetUserByID(ctx, id)
	}

	user, err := s.cacheStorage.Users.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		user, err = s.store.Users.GetUserByID(ctx, id)
		if err != nil {
			return nil, err
		}

		if err := s.cacheStorage.Users.Set(ctx, user); err != nil {
			return nil, err
		}
	}

	return user, nil
}
