package application

import (
	"jaehonam.com/lmnop/application/port/in"
	"jaehonam.com/lmnop/application/port/out"
)

// Check Service implements in.Query
var _ = in.Query(&Service{})

type Service struct {
	*UserService
	*ProblemService
}

func NewService(client out.Port) *Service {
	return &Service{
		UserService:    NewUserService(client),
		ProblemService: NewProblemService(client),
	}
}
