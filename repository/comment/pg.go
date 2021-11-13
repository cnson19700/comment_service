package comment

import (
	"context"
	"strings"

	bookmodel "github.com/cnson19700/book_service/model"
	"github.com/cnson19700/comment_service/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type pgRepository struct {
	getClient func(ctx context.Context) *gorm.DB
}

func NewPGRepository(getClient func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getClient}
}

func (r *pgRepository) GetById(ctx context.Context, ID int64) (*model.Comment, error) {
	db := r.getClient(ctx)
	comment := &model.Comment{}

	err := db.Where("id = ?", ID).
		First(comment).Error

	if err != nil {
		return nil, errors.Wrap(err, "get comment by id")
	}

	return comment, nil
}

func (r *pgRepository) GetAll(ctx context.Context) ([]model.Comment, error) {
	db := r.getClient(ctx)
	listComment := []model.Comment{}

	db.Find(&listComment)

	return listComment, nil
}

func (r *pgRepository) Insert(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	db := r.getClient(ctx)
	err := db.Create(comment).Error

	return comment, errors.Wrap(err, "create comment")
}

func (r *pgRepository) Delete(ctx context.Context, id int64) error {
	db := r.getClient(ctx)
	err := db.Where("id = ?", id).Delete(&model.Comment{}).Error

	return errors.Wrap(err, "delete comment fail")
}

func (r *pgRepository) DeleteSubComment(ctx context.Context, parentID int64) error {
	db := r.getClient(ctx)
	err := db.Where("id = ?", parentID).Delete(&model.Comment{}).Error

	return errors.Wrap(err, "delete comment fail")
}

func (r *pgRepository) Update(ctx context.Context, book *model.Comment) (*model.Comment, error) {
	db := r.getClient(ctx)
	err := db.Save(book).Error

	return book, errors.Wrap(err, "update comment fail")
}

func (r *pgRepository) Find(
	ctx context.Context,
	conditions []bookmodel.Condition,
	paginator *bookmodel.Paginator,
	orders []string,
) (*model.CommentResult, error) {
	// Build query
	db := r.getClient(ctx)
	query := db.Model(&model.Comment{})

	// Where
	for _, condition := range conditions {
		switch strings.ToLower(condition.Type) {
		case bookmodel.ConditionTypeNot:
			query.Not(condition.Pattern, condition.Values...)
		case bookmodel.ConditionTypeOr:
			query.Or(condition.Pattern, condition.Values...)
		default:
			query.Where(condition.Pattern, condition.Values...)
		}
	}

	// Order
	for _, order := range orders {
		query.Order(order)
	}

	// Paging
	var result model.CommentResult

	if paginator.Limit >= 0 {
		if paginator.Page <= 0 {
			paginator.Page = 1
		}

		if paginator.Limit == 0 {
			paginator.Limit = bookmodel.PageSize
		}

		result.Page = paginator.Page
		result.Limit = paginator.Limit
		query.Count(&result.Total).Scopes(paginator.Paginate())
	}

	err := query.Find(&result.Data).Error

	return &result, err
}
