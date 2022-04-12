package in

import "jaehonam.com/lmnop/domain"

type UserQuery interface {
	GetUser(userID string) (*domain.User, error)
}
