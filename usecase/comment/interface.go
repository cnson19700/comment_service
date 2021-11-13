package comment

import (
	"context"

	"github.com/soncaodb/model"
)

type IUsecase interface {
	Insert(ctx context.Context, req *InsertCommentRequest) (*model.Comment, error)
	Delete(ctx context.Context, req DeleteCommentRequest) error
	GetList(ctx context.Context, req *GetListRequest) (*model.CommentResult, error)
}
