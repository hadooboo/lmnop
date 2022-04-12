package in

import "jaehonam.com/lmnop/domain"

type ProblemQuery interface {
	GetOptimumProblem(userID string, tier domain.Tier, except []int64) (*domain.Problem, error)
}
