package comment

import (
	"context"

	"github.com/soncaodb/model"
)

type Repository interface {
	GetById(ctx context.Context, ID int64) (*model.Comment, error)
	GetAll(ctx context.Context) ([]model.Comment, error)
	Insert(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	Delete(ctx context.Context, ID int64) error
	DeleteSubComment(ctx context.Context, parentID int64) error
	Update(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	Find(
		ctx context.Context,
		conditions []model.Condition,
		paginator *model.Paginator,
		orders []string,
	) (*model.CommentResult, error)
}
