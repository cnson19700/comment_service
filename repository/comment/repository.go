package comment

import (
	"context"

	bookmodel "github.com/cnson19700/book_service/model"
	"github.com/cnson19700/comment_service/model"
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
		conditions []bookmodel.Condition,
		paginator *bookmodel.Paginator,
		orders []string,
	) (*model.CommentResult, error)
}
