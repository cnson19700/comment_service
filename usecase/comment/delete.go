package comment

import (
	"context"
	"log"
)

type DeleteCommentRequest struct {
	ID int64 `json:"id"`
}

func (u *Usecase) Delete(ctx context.Context, req DeleteCommentRequest) error {
	// claim := middleware.GetClaim(ctx)
	// userID := claim.UserID

	// user, err := u.userRepo.GetById(ctx, userID)
	// if err != nil {
	// 	return myerror.ErrGetUser(err)
	// }

	comment, err := u.commentRepo.GetById(ctx, req.ID)
	if err != nil {
		log.Fatal(err)
	}

	// if user.Role == "user" && comment.UserID != userID {
	// 	return myerror.ErrUserComment(err)
	// }

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
