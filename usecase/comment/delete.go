package comment

import (
	"context"
)

type DeleteCommentRequest struct {
	ID int64 `json:"id"`
}

func (u *Usecase) Delete(ctx context.Context, req DeleteCommentRequest) error {

	comment, err := u.commentRepo.GetById(ctx, req.ID)
	if err != nil {
		return err
	}

	// check if this is root comment and delete subcomment
	if comment.ParentID != 1 {
		err := u.commentRepo.DeleteSubComment(ctx, comment.ParentID)

		if err != nil {
			return err
		}
	}

	err = u.commentRepo.Delete(ctx, comment.ID)
	if err != nil {
		return err
	}

	return err
}
