package gin_server_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	gin_server "jaehonam.com/lmnop/adapter/in/gin-server"
	"jaehonam.com/lmnop/application/port/in"
	"jaehonam.com/lmnop/domain"
)

var (
	testUserID = "njh21598"
	testUser   = &domain.User{
		Handle:          testUserID,
		SolvedCount:     500,
		Tier:            domain.Gold5Tier,
		Rating:          1500,
		Class:           4,
		ClassDecoration: domain.Gold,
		MaxStreak:       10,
		Top100: []*domain.Problem{
			{ProblemID: 1, Level: domain.Gold1Level},
			{ProblemID: 2, Level: domain.Gold2Level},
			{ProblemID: 3, Level: domain.Gold5Level},
			{ProblemID: 4, Level: domain.Silver2Level},
			{ProblemID: 5, Level: domain.Bronze2Level},
		},
	}
	testProblem = &domain.Problem{
		ProblemID: 6,
		Level:     domain.Platinum5Level,
	}
)

// Check mockQuery implements in.Query
var _ = in.Query(&mockQuery{})

type mockQuery struct {
}

func (r *mockQuery) GetUser(userID string) (*domain.User, error) {
	return testUser, nil
}

func (r *mockQuery) GetOptimumProblem(userID string, level domain.Level, except []int64) (*domain.Problem, error) {
	return testProblem, nil
}

var (
	server *gin_server.GinAPIServer
)

func setUp() {
	server = gin_server.NewGinAPIServer(&mockQuery{}, 8080)
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func tearDown() {
	server = nil
}

func TestGetUserSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/njh21598", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetOptimumProblemSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/optimum-problem?handle=njh21598&level=12&except=1000", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetOptimumProblemFailNoHandle(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/optimum-problem?level=12", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetOptimumProblemFailNoLevel(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/optimum-problem?handle=njh21598", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetOptimumProblemFailInvalidLevel(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/optimum-problem?handle=njh21598&level=abc", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetOptimumProblemFailInvalidExcept(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/optimum-problem?handle=njh21598&level=12&except=1000,abc", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
