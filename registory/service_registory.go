package registory

import (
	"github.com/trewanek/layered-architecture/domain/service"
)

type Service interface {
	NewUserService() service.UserService
}

type serviceImpl struct {
	rep     Repository
	userSer service.UserService
}

func NewService(rep Repository) Service {
	return &serviceImpl{
		rep: rep,
	}
}

func (s *serviceImpl) NewUserService() service.UserService {
	if s.userSer == nil {
		s.userSer = service.NewUserService(s.rep.NewUserRepository())
	}
	return s.userSer
}
