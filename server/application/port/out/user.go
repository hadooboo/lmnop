package out

import "jaehonam.com/lmnop/domain"

type UserPort interface {
	GetUser(userID string) (*domain.UserDTO, error)

	GetUserProblemStats(userID string) (*domain.UserProblemStatsDTO, error)

	GetUserProblemTagStats(userID string) (*domain.UserProblemTagStatsDTO, error)

	GetUserTop100(userID string) (*domain.UserTop100DTO, error)
}
