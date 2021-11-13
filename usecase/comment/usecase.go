package comment

import (
	"github.com/cnson19700/comment_service/repository"
	"github.com/cnson19700/comment_service/repository/comment"
	userrepository "github.com/cnson19700/user_service/repository/user"
)

type Usecase struct {
	commentRepo comment.Repository
	userRepo    userrepository.Repository
}

func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		commentRepo: repo.Comment,
		userRepo:    repo.User,
	}
}
