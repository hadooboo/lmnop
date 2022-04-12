package out

import "jaehonam.com/lmnop/domain"

type ProblemPort interface {
	GetProblemsByTier(level domain.Level, page int) (*domain.ProblemsDTO, error)

	GetProblemsByTierAndSolvedBy(level domain.Level, solvedBy string, page int) (*domain.ProblemsDTO, error)
}
