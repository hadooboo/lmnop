package solved_ac_client

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"jaehonam.com/lmnop/application/port/out"
	"jaehonam.com/lmnop/domain"
)

const (
	baseURL = "https://solved.ac/api/v3"

	defaultRetryCount = 3

	pathGetUser             = "/user/show"
	pathGetUserProblemStats = "/user/problem_stats"
)

type RestyAdapter struct {
	client *resty.Client
}

// Check RestyAdapter implements out.Port
var _ = out.Port(&RestyAdapter{})

func NewRestyAdapter() *RestyAdapter {
	return &RestyAdapter{
		client: resty.New().
			SetBaseURL(baseURL).
			SetRetryCount(defaultRetryCount),
	}
}

func (r *RestyAdapter) GetUser(userID string) (*domain.User, error) {
	var result userResp
	resp, err := r.client.R().
		SetQueryString(fmt.Sprintf("handle=%v", userID)).
		ForceContentType("application/json").
		SetResult(&result).
		Get(pathGetUser)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode() {
	case http.StatusOK:
		return result.mapToDomainEntity(), nil
	case http.StatusNotFound:
		return nil, domain.ErrNotFoundUser
	default:
		return nil, domain.ErrUnexpected
	}
}

func (r *RestyAdapter) GetUserProblemStats(userID string) (*domain.UserProblemStats, error) {
	var result userProblemStatsResp
	resp, err := r.client.R().
		SetQueryString(fmt.Sprintf("handle=%v", userID)).
		ForceContentType("application/json").
		SetResult(&result).
		Get(pathGetUserProblemStats)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode() {
	case http.StatusOK:
		return result.mapToDomainEntity(), nil
	case http.StatusNotFound:
		return nil, domain.ErrNotFoundUser
	default:
		return nil, domain.ErrUnexpected
	}
}
