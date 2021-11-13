package comment

import (
	"context"

	"github.com/cnson19700/comment_service/model"
	checkform "github.com/cnson19700/comment_service/package/checkForm"
	"github.com/cnson19700/comment_service/util/myerror"
	"github.com/cnson19700/pkg/middleware"
)

type InsertCommentRequest struct {
	BookID   int64  `json:"book_id"`
	ParentID int64  `json:"parent_id"`
	Content  string `json:"content"`
}

func (u *Usecase) Insert(ctx context.Context, req *InsertCommentRequest) (*model.Comment, error) {
	comment := &model.Comment{}

	//Get current user from Token
	payload := middleware.GetClaim(ctx)
	UserID := payload.UserID

	//Check content format
	isAllow, content := checkform.CheckFormatValue("content", req.Content)

	if !isAllow {
		return nil, myerror.ErrContentFormat(nil)
	}

	comment.Content = content

	//check if insert subcmt
	if req.ParentID != 0 {
		comment.ParentID = req.ParentID
	}

	comment.BookID = req.BookID
	comment.UserID = UserID

	res, err := u.commentRepo.Insert(ctx, comment)

	if err != nil {
		return nil, err
	}
	return res, nil
}
