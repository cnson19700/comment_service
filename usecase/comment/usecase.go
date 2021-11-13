package comment

import (
	"github.com/soncaodb/repository"
	"github.com/soncaodb/repository/comment"
	"github.com/soncaodb/repository/user"
)

type Usecase struct {
	commentRepo comment.Repository
	userRepo    user.Repository
}

func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		commentRepo: repo.Comment,
		userRepo:    repo.User,
	}
}
