package application_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"jaehonam.com/lmnop/application"
	"jaehonam.com/lmnop/application/port/out"
	"jaehonam.com/lmnop/domain"
)

var (
	testUserID  = "njh21598"
	testUserDTO = &domain.UserDTO{
		Handle:              testUserID,
		Bio:                 "test user",
		SolvedCount:         500,
		VoteCount:           100,
		Exp:                 10001000,
		Tier:                domain.Gold5Tier,
		Rating:              1500,
		RatingByProblemsSum: 1200,
		RatingByClass:       150,
		RatingBySolvedCount: 150,
		RatingByVoteCount:   100,
		Class:               4,
		ClassDecoration:     domain.Gold,
		RivalCount:          0,
		ReverseRivalCount:   0,
		MaxStreak:           10,
		Rank:                4000,
	}
	testUserTop100DTO = &domain.UserTop100DTO{
		Count: 5,
		Items: []*domain.ProblemDTO{
			{ProblemID: 1, Level: domain.Gold1Level},
			{ProblemID: 2, Level: domain.Gold2Level},
			{ProblemID: 3, Level: domain.Gold5Level},
			{ProblemID: 4, Level: domain.Silver2Level},
			{ProblemID: 5, Level: domain.Bronze2Level},
		},
	}
	testProblemsDTOAllPartial = []*domain.ProblemsDTO{
		{
			Count: 8,
			Items: []*domain.ProblemDTO{
				{ProblemID: 1, Level: domain.Gold4Level},
				{ProblemID: 2, Level: domain.Gold4Level},
				{ProblemID: 3, Level: domain.Gold4Level},
			},
		},
		{
			Count: 8,
			Items: []*domain.ProblemDTO{
				{ProblemID: 4, Level: domain.Gold4Level},
				{ProblemID: 5, Level: domain.Gold4Level},
				{ProblemID: 6, Level: domain.Gold4Level},
			},
		},
		{
			Count: 8,
			Items: []*domain.ProblemDTO{
				{ProblemID: 7, Level: domain.Gold4Level},
				{ProblemID: 8, Level: domain.Gold4Level},
			},
		},
		{
			Count: 8,
			Items: []*domain.ProblemDTO{},
		},
	}
	testProblemsDTOSolvedPartial = []*domain.ProblemsDTO{
		{
			Count: 4,
			Items: []*domain.ProblemDTO{
				{ProblemID: 1, Level: domain.Gold4Level},
				{ProblemID: 3, Level: domain.Gold4Level},
				{ProblemID: 6, Level: domain.Gold4Level},
			},
		},
		{
			Count: 4,
			Items: []*domain.ProblemDTO{
				{ProblemID: 7, Level: domain.Gold4Level},
			},
		},
		{
			Count: 4,
			Items: []*domain.ProblemDTO{},
		},
	}
	testProblemsDTOAllAll = []*domain.ProblemsDTO{
		{
			Count: 2,
			Items: []*domain.ProblemDTO{
				{ProblemID: 1, Level: domain.Gold3Level},
				{ProblemID: 2, Level: domain.Gold3Level},
			},
		},
		{
			Count: 2,
			Items: []*domain.ProblemDTO{},
		},
	}
	testProblemsDTOSolvedAll = []*domain.ProblemsDTO{
		{
			Count: 2,
			Items: []*domain.ProblemDTO{
				{ProblemID: 1, Level: domain.Gold3Level},
				{ProblemID: 2, Level: domain.Gold3Level},
			},
		},
		{
			Count: 2,
			Items: []*domain.ProblemDTO{},
		},
	}
	testProblemsDTOAllNone = []*domain.ProblemsDTO{
		{
			Count: 2,
			Items: []*domain.ProblemDTO{
				{ProblemID: 1, Level: domain.Gold2Level},
				{ProblemID: 2, Level: domain.Gold2Level},
			},
		},
		{
			Count: 2,
			Items: []*domain.ProblemDTO{},
		},
	}
	testProblemsDTOSolvedNone = []*domain.ProblemsDTO{
		{
			Count: 0,
			Items: []*domain.ProblemDTO{},
		},
	}
)

// Check mockPort implements out.Port
var _ = out.Port(&mockPort{})

type mockPort struct {
}

func (r *mockPort) GetUser(userID string) (*domain.UserDTO, error) {
	return testUserDTO, nil
}

func (r *mockPort) GetUserProblemStats(userID string) (*domain.UserProblemStatsDTO, error) {
	panic("not implemented")
}

func (r *mockPort) GetUserProblemTagStats(userID string) (*domain.UserProblemTagStatsDTO, error) {
	panic("not implemented")
}

func (r *mockPort) GetUserTop100(userID string) (*domain.UserTop100DTO, error) {
	return testUserTop100DTO, nil
}

func (r *mockPort) GetProblemsByTier(level domain.Level, page int) (*domain.ProblemsDTO, error) {
	switch level {
	case domain.Gold4Level:
		return testProblemsDTOAllPartial[page-1], nil
	case domain.Gold3Level:
		return testProblemsDTOAllAll[page-1], nil
	case domain.Gold2Level:
		return testProblemsDTOAllNone[page-1], nil
	}
	panic("not implemented")
}

func (r *mockPort) GetProblemsByTierAndSolvedBy(level domain.Level, solvedBy string, page int) (*domain.ProblemsDTO, error) {
	switch level {
	case domain.Gold4Level:
		return testProblemsDTOSolvedPartial[page-1], nil
	case domain.Gold3Level:
		return testProblemsDTOSolvedAll[page-1], nil
	case domain.Gold2Level:
		return testProblemsDTOSolvedNone[page-1], nil
	}
	panic("not implemented")
}

var (
	service *application.Service
)

func setUp() {
	service = application.NewService(&mockPort{})
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func tearDown() {
	service = nil
}

func TestGetUser(t *testing.T) {
	user, err := service.GetUser(testUserID)
	assert.NoError(t, err)
	assert.Equal(t, user.Handle, testUserID)
	assert.Equal(t, user.SolvedCount, testUserDTO.SolvedCount)
	assert.Equal(t, user.Tier, testUserDTO.Tier)
	assert.Equal(t, user.Rating, testUserDTO.Rating)
	assert.Equal(t, user.Class, testUserDTO.Class)
	assert.Equal(t, user.ClassDecoration, testUserDTO.ClassDecoration)
	assert.Equal(t, user.MaxStreak, testUserDTO.MaxStreak)

	assert.NotNil(t, user.Top100)
	assert.Equal(t, len(user.Top100), len(testUserTop100DTO.Items))
	for i := range user.Top100 {
		assert.Equal(t, user.Top100[i].ProblemID, testUserTop100DTO.Items[i].ProblemID)
		assert.Equal(t, user.Top100[i].Level, testUserTop100DTO.Items[i].Level)
	}
}

func TestGetOptimumProblem(t *testing.T) {
	problem, err := service.GetOptimumProblem(testUserID, domain.Gold4Level, []int64{})
	assert.NoError(t, err)
	assert.NotNil(t, problem)
	assert.EqualValues(t, problem.ProblemID, 2)

	problem, err = service.GetOptimumProblem(testUserID, domain.Gold4Level, []int64{2})
	assert.NoError(t, err)
	assert.NotNil(t, problem)
	assert.EqualValues(t, problem.ProblemID, 4)

	problem, err = service.GetOptimumProblem(testUserID, domain.Gold4Level, []int64{2, 4})
	assert.NoError(t, err)
	assert.NotNil(t, problem)
	assert.EqualValues(t, problem.ProblemID, 5)

	problem, err = service.GetOptimumProblem(testUserID, domain.Gold4Level, []int64{2, 4, 5})
	assert.NoError(t, err)
	assert.NotNil(t, problem)
	assert.EqualValues(t, problem.ProblemID, 8)

	problem, err = service.GetOptimumProblem(testUserID, domain.Gold4Level, []int64{2, 4, 5, 8})
	assert.Error(t, err, domain.ErrNotFoundProblem)

	problem, err = service.GetOptimumProblem(testUserID, domain.Gold3Level, []int64{})
	assert.Error(t, err, domain.ErrNotFoundProblem)

	problem, err = service.GetOptimumProblem(testUserID, domain.Gold2Level, []int64{})
	assert.NoError(t, err)
	assert.NotNil(t, problem)
	assert.EqualValues(t, problem.ProblemID, 1)

	problem, err = service.GetOptimumProblem(testUserID, domain.Gold2Level, []int64{1})
	assert.NoError(t, err)
	assert.NotNil(t, problem)
	assert.EqualValues(t, problem.ProblemID, 2)

	problem, err = service.GetOptimumProblem(testUserID, domain.Gold2Level, []int64{1, 2})
	assert.Error(t, err, domain.ErrNotFoundProblem)
}
