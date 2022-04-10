package out

import "jaehonam.com/lmnop/domain"

type UserPort interface {
	GetUser(userID string) (*domain.User, error)

	GetUserProblemStats(userID string) (*domain.UserProblemStats, error)
}
