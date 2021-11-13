package repository

import (
	"context"

	"github.com/cnson19700/book_service/repository/book"
	"github.com/cnson19700/user_service/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	User user.Repository
	Book book.Repository
}

func New(
	getSQLClient func(ctx context.Context) *gorm.DB,
	// getRedisClient func(ctx context.Context) *redis.Client,
) *Repository {
	return &Repository{
		User: user.NewPGRepository(getSQLClient),
	}
}
