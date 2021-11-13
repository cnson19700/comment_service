package repository

import (
	"context"

	"github.com/cnson19700/comment_service/repository/comment"
	"github.com/cnson19700/user_service/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	User    user.Repository
	Comment comment.Repository
}

func New(
	getSQLClient func(ctx context.Context) *gorm.DB,
	// getRedisClient func(ctx context.Context) *redis.Client,
) *Repository {
	return &Repository{
		User:    user.NewPGRepository(getSQLClient),
		Comment: comment.NewPGRepository(getSQLClient),
	}
}
