package domain

type User struct {
	Handle              string
	Bio                 string
	SolvedCount         int64
	VoteCount           int64
	Exp                 int64
	Tier                Tier
	Rating              int64
	RatingByProblemsSum int64
	RatingByClass       int64
	RatingBySolvedCount int64
	RatingByVoteCount   int64
	Class               int64
	ClassDecoration     ClassDecoration
	RivalCount          int64
	ReverseRivalCount   int64
	MaxStreak           int64
	Rank                int64
}

type ClassDecoration string

const (
	None   ClassDecoration = "none"
	Silver ClassDecoration = "silver"
	Gold   ClassDecoration = "gold"
)

type UserProblemStats []*UserProblemStat

type UserProblemStat struct {
	Level   Level
	Total   int64
	Solved  int64
	Partial int64
	Tried   int64
	Exp     int64
}

type UserProblemTagStats struct {
	Count int64
	Items []*UserProblemTagStat
}

type UserProblemTagStat struct {
	Tag     string
	Total   int64
	Solved  int64
	Partial int64
	Tried   int64
	Exp     int64
}

type UserTop100 struct {
	Count int64
	Items []*Problem
}
