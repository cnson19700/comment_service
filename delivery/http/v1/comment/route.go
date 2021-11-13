package comment

import (
	"github.com/labstack/echo/v4"

	"github.com/cnson19700/comment_service/usecase"
	"github.com/cnson19700/comment_service/usecase/comment"
)

type Route struct {
	commentUseCase comment.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		commentUseCase: useCase.Comment,
	}
	group.POST("", r.Insert)
	group.DELETE("/:id", r.Delete)
	group.GET("", r.GetList)
}
