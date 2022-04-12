package application

import (
	"jaehonam.com/lmnop/application/port/out"
	"jaehonam.com/lmnop/domain"
)

type ProblemService struct {
	client out.ProblemPort
}

func NewProblemService(client out.ProblemPort) *ProblemService {
	return &ProblemService{
		client: client,
	}
}

func (r *ProblemService) GetOptimumProblem(userID string, tier domain.Tier, except []int64) (*domain.Problem, error) {
	panic("not implemented")
}
