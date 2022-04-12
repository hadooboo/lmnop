package out

import "jaehonam.com/lmnop/domain"

type SearchPort interface {
	GetProblemsByTier(tier domain.Tier, page int) (*domain.ProblemsDTO, error)

	GetProblemsByTierAndSolvedBy(tier domain.Tier, solvedBy string, page int) (*domain.ProblemsDTO, error)
}
