package comment

import (
	"errors"
	"strconv"

	"github.com/cnson19700/comment_service/model"
	"github.com/cnson19700/comment_service/usecase/comment"
	"github.com/cnson19700/pkg/apperror"
	"github.com/cnson19700/pkg/utils"
	"github.com/labstack/echo/v4"
)

func (r *Route) GetList(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
	)

	bookID, _ := strconv.Atoi(c.QueryParam("book_id"))
	parentID, _ := strconv.Atoi(c.QueryParam("parent_id"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	paginator := model.Paginator{
		Page:  page,
		Limit: limit,
	}

	filter := model.CommentFilter{
		BookID:   int64(bookID),
		ParentID: int64(parentID),
	}

	if err := c.Bind(&paginator); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	if err := c.Bind(&filter); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	req := comment.GetListRequest{
		Filter:    &filter,
		Paginator: &paginator,
	}

	res, err := r.commentUseCase.GetList(ctx, &req)
	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}
	return utils.Response.Success(c, res)
}
