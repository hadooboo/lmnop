package in

import "jaehonam.com/lmnop/domain"

type ProblemQuery interface {
	GetOptimumProblem(userID string, level domain.Level, except []int64) (*domain.Problem, error)
}
