package gin_server

import (
	"encoding/json"

	"jaehonam.com/lmnop/domain"
)

type getUserResp struct {
	user *domain.User
}

func (r *getUserResp) MarshalJSON() ([]byte, error) {
	l := make([]*problemResp, len(r.user.Top100))
	for i, v := range r.user.Top100 {
		l[i] = &problemResp{
			ProblemID: v.ProblemID,
			Level:     int(v.Level),
		}
	}
	v := &userResp{
		Handle:          r.user.Handle,
		SolvedCount:     r.user.SolvedCount,
		Tier:            int(r.user.Tier),
		Rating:          r.user.Rating,
		Class:           r.user.Class,
		ClassDecoration: string(r.user.ClassDecoration),
		MaxStreak:       r.user.MaxStreak,
		Top100:          l,
	}
	return json.Marshal(v)
}

type userResp struct {
	Handle          string         `json:"handle"`
	SolvedCount     int64          `json:"solved_count"`
	Tier            int            `json:"tier"`
	Rating          int64          `json:"rating"`
	Class           int64          `json:"class"`
	ClassDecoration string         `json:"class_decoration"`
	MaxStreak       int64          `json:"max_streak"`
	Top100          []*problemResp `json:"top_100"`
}

type problemResp struct {
	ProblemID int64 `json:"problem_id"`
	Level     int   `json:"level"`
}

type getOptimumProblemResp struct {
	problem *domain.Problem
}

func (r *getOptimumProblemResp) MarshalJSON() ([]byte, error) {
	v := &problemResp{
		ProblemID: r.problem.ProblemID,
		Level:     int(r.problem.Level),
	}
	return json.Marshal(v)
}
