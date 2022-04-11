package solved_ac_client_test

import (
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"jaehonam.com/lmnop/adapter/out/solved-ac-client"
	"jaehonam.com/lmnop/domain"
)

var (
	adapter *solved_ac_client.RestyAdapter
)

func setUp() {
	adapter = solved_ac_client.NewRestyAdapter()
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func tearDown() {
	adapter = nil
}

func TestGetUserSuccess(t *testing.T) {
	userID := "njh21598"
	user, err := adapter.GetUser(userID)
	assert.NoError(t, err)
	assert.Equal(t, user.Handle, userID)
	t.Log(spew.Sdump(user))
}

func TestGetUserFailNotFoundUser(t *testing.T) {
	userID := ""
	_, err := adapter.GetUser(userID)
	assert.Equal(t, err, domain.ErrNotFoundUser)
}

func TestGetUserProblemStatsSuccess(t *testing.T) {
	userID := "njh21598"
	userProblemStats, err := adapter.GetUserProblemStats(userID)
	assert.NoError(t, err)
	t.Log(spew.Sdump(userProblemStats))
}

func TestGetUserProblemStatsFailNotFoundUser(t *testing.T) {
	userID := ""
	_, err := adapter.GetUserProblemStats(userID)
	assert.Equal(t, err, domain.ErrNotFoundUser)
}

func TestGetUserProblemTagStatsSuccess(t *testing.T) {
	userID := "njh21598"
	userProblemTagStats, err := adapter.GetUserProblemTagStats(userID)
	assert.NoError(t, err)
	t.Log(spew.Sdump(userProblemTagStats))
}

func TestGetUserProblemTagStatsFailNotFoundUser(t *testing.T) {
	userID := ""
	_, err := adapter.GetUserProblemTagStats(userID)
	assert.Equal(t, err, domain.ErrNotFoundUser)
}

func TestGetUserTop100Success(t *testing.T) {
	userID := "njh21598"
	userTop100, err := adapter.GetUserTop100(userID)
	assert.NoError(t, err)
	t.Log(spew.Sdump(userTop100))
}

func TestGetUserTop100FailNotFoundUser(t *testing.T) {
	userID := ""
	_, err := adapter.GetUserTop100(userID)
	assert.Equal(t, err, domain.ErrNotFoundUser)
}
