package comment

import (
	"context"
	"fmt"

	bookmodel "github.com/cnson19700/book_service/model"
	"github.com/cnson19700/comment_service/model"
	"github.com/cnson19700/comment_service/util/myerror"
)

type GetListRequest struct {
	Filter    *model.CommentFilter
	Paginator *bookmodel.Paginator
	OrderBy   string `json:"order_by,omitempty" query:"order_by"`
	OrderType string `json:"order_type,omitempty" query:"order_type"`
}

func (u *Usecase) GetList(ctx context.Context, req *GetListRequest) (*model.CommentResult, error) {
	listBookID := make([]interface{}, 1)
	listBookID[0] = req.Filter.BookID

	listParentID := make([]interface{}, 1)
	listParentID[0] = 1 //Default 1 is root comment, not subcomment

	if req.Filter.ParentID != 0 {
		listParentID[0] = req.Filter.ParentID
	}

	conditions := []bookmodel.Condition{
		{Pattern: "book_id",
			Values: listBookID},
		{Pattern: "parent_id",
			Values: listParentID},
	}

	//Order
	orders := make([]string, 0)
	if req.OrderBy != "" {
		orders = []string{fmt.Sprintf("%s %s", req.OrderBy, req.OrderType)}
	}

	//Paging
	paginator := &bookmodel.Paginator{
		Page:  1,
		Limit: 20,
	}

	if req.Paginator != nil {
		paginator = req.Paginator
	}

	res, err := u.commentRepo.Find(ctx, conditions, paginator, orders)
	if err != nil {
		return nil, myerror.ErrFindComment(err)
	}

	return res, nil
}
