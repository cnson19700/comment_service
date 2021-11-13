package usecase

import (
	"github.com/cnson19700/book_service/repository"
	"github.com/cnson19700/book_service/usecase/book"
)

type UseCase struct {
	Book book.IUsecase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		Book: book.New(repo),
	}
}
