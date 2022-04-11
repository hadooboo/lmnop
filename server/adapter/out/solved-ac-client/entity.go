package solved_ac_client

type userResp struct {
	Handle              string `json:"handle"`
	Bio                 string `json:"bio"`
	SolvedCount         int64  `json:"solvedCount"`
	VoteCount           int64  `json:"voteCount"`
	Exp                 int64  `json:"exp"`
	Tier                int    `json:"tier"`
	Rating              int64  `json:"rating"`
	RatingByProblemsSum int64  `json:"ratingByProblemsSum"`
	RatingByClass       int64  `json:"ratingByClass"`
	RatingBySolvedCount int64  `json:"ratingBySolvedCount"`
	RatingByVoteCount   int64  `json:"ratingByVoteCount"`
	Class               int64  `json:"class"`
	ClassDecoration     string `json:"classDecoration"`
	RivalCount          int64  `json:"rivalCount"`
	ReverseRivalCount   int64  `json:"reverseRivalCount"`
	MaxStreak           int64  `json:"maxStreak"`
	Rank                int64  `json:"rank"`
}

type userProblemStatsResp []*userProblemStatResp

type userProblemStatResp struct {
	Level   int   `json:"level"`
	Total   int64 `json:"total"`
	Solved  int64 `json:"solved"`
	Partial int64 `json:"partial"`
	Tried   int64 `json:"tried"`
	Exp     int64 `json:"exp"`
}

type userProblemTagStatsResp struct {
	Count int64 `json:"count"`
	Items []*userProblemTagStatResp
}

type userProblemTagStatResp struct {
	Tag     *tagResp `json:"tag"`
	Total   int64    `json:"total"`
	Solved  int64    `json:"solved"`
	Partial int64    `json:"partial"`
	Tried   int64    `json:"tried"`
	Exp     int64    `json:"exp"`
}

type tagResp struct {
	Key string `json:"key"`
}

type userTop100Resp struct {
	Count int64          `json:"count"`
	Items []*problemResp `json:"items"`
}

type problemResp struct {
	ProblemID int64 `json:"problemId"`
	Level     int   `json:"level"`
}
