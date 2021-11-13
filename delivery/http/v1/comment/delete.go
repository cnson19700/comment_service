package comment

import (
	"errors"
	"strconv"

	"github.com/cnson19700/comment_service/usecase/comment"
	"github.com/cnson19700/pkg/apperror"
	"github.com/cnson19700/pkg/utils"
	"github.com/labstack/echo/v4"
)

func (r *Route) Delete(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
	)

	commentID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	// Bind order by
	if err := c.Bind(&comment.DeleteCommentRequest{}); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	err := r.commentUseCase.Delete(ctx, comment.DeleteCommentRequest{
		ID: commentID,
	})

	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}

	return utils.Response.Success(ctx, nil)
}
