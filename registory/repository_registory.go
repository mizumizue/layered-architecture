package registory

import (
	"github.com/trewanek/layered-architecture/domain/repository"
	"github.com/trewanek/layered-architecture/infrastructure/persistence/firestore"
)

type Repository interface {
	NewUserRepository() repository.UserRepository
}

type repositoryImpl struct {
	userRepo repository.UserRepository
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) NewUserRepository() repository.UserRepository {
	if r.userRepo == nil {
		r.userRepo = firestore.NewUserRepository()
	}
	return r.userRepo
}
