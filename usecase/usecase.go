package usecase

import (
	"github.com/cnson19700/comment_service/repository"
	"github.com/cnson19700/comment_service/usecase/comment"
)

type UseCase struct {
	Comment comment.IUsecase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		Comment: comment.New(repo),
	}
}
