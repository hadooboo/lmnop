package application

import (
	"jaehonam.com/lmnop/application/port/out"
	"jaehonam.com/lmnop/domain"
)

type UserService struct {
	client out.UserPort
}

func NewUserService(client out.UserPort) *UserService {
	return &UserService{
		client: client,
	}
}

func (r *UserService) GetUser(userID string) (*domain.User, error) {
	userDTO, err := r.client.GetUser(userID)
	if err != nil {
		return nil, err
	}

	userTop100DTO, err := r.client.GetUserTop100(userID)
	if err != nil {
		return nil, err
	}

	top100 := make([]*domain.Problem, len(userTop100DTO.Items))
	for i, v := range userTop100DTO.Items {
		top100[i] = &domain.Problem{
			ProblemID: v.ProblemID,
			Level:     v.Level,
		}
	}

	return &domain.User{
		Handle:          userDTO.Handle,
		SolvedCount:     userDTO.SolvedCount,
		Tier:            userDTO.Tier,
		Rating:          userDTO.Rating,
		Class:           userDTO.Class,
		ClassDecoration: userDTO.ClassDecoration,
		MaxStreak:       userDTO.MaxStreak,
		Top100:          top100,
	}, nil
}
