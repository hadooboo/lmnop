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

	pathGetUser                = "/user/show"
	pathGetUserProblemStats    = "/user/problem_stats"
	pathGetUserProblemTagStats = "/user/problem_tag_stats"
	pathGetUserTop100          = "/user/top_100"
	pathSearchProblem          = "/search/problem"
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

func (r *RestyAdapter) GetUserProblemTagStats(userID string) (*domain.UserProblemTagStats, error) {
	var result userProblemTagStatsResp
	resp, err := r.client.R().
		SetQueryString(fmt.Sprintf("handle=%v", userID)).
		ForceContentType("application/json").
		SetResult(&result).
		Get(pathGetUserProblemTagStats)
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

func (r *RestyAdapter) GetUserTop100(userID string) (*domain.UserTop100, error) {
	var result userTop100Resp
	resp, err := r.client.R().
		SetQueryString(fmt.Sprintf("handle=%v", userID)).
		ForceContentType("application/json").
		SetResult(&result).
		Get(pathGetUserTop100)
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

func (r *RestyAdapter) GetProblemsByTier(tier domain.Tier, page int) (*domain.Problems, error) {
	var result problemsByTierResp
	resp, err := r.client.R().
		SetQueryString(fmt.Sprintf("query=solvable:true+tier:%v&page=%v&sort=solved&direction=desc", tier, page)).
		ForceContentType("application/json").
		SetResult(&result).
		Get(pathSearchProblem)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode() {
	case http.StatusOK:
		return result.mapToDomainEntity(), nil
	default:
		return nil, domain.ErrUnexpected
	}
}

func (r *RestyAdapter) GetProblemsByTierAndSolvedBy(tier domain.Tier, solvedBy string, page int) (*domain.Problems, error) {
	var result problemsByTierAndSolvedByResp
	resp, err := r.client.R().
		SetQueryString(fmt.Sprintf("query=solvable:true+tier:%v+solved_by:%v&page=%v&sort=solved&direction=desc", tier, solvedBy, page)).
		ForceContentType("application/json").
		SetResult(&result).
		Get(pathSearchProblem)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode() {
	case http.StatusOK:
		return result.mapToDomainEntity(), nil
	default:
		return nil, domain.ErrUnexpected
	}
}
