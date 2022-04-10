package solved_ac_client

import "jaehonam.com/lmnop/domain"

func (r userResp) mapToDomainEntity() *domain.User {
	return &domain.User{
		Handle:              r.Handle,
		Bio:                 r.Bio,
		SolvedCount:         r.SolvedCount,
		VoteCount:           r.VoteCount,
		Exp:                 r.Exp,
		Tier:                domain.Tier(r.Tier),
		Rating:              r.Rating,
		RatingByProblemsSum: r.RatingByProblemsSum,
		RatingByClass:       r.RatingByClass,
		RatingBySolvedCount: r.RatingBySolvedCount,
		RatingByVoteCount:   r.RatingByVoteCount,
		Class:               r.Class,
		ClassDecoration:     domain.ClassDecoration(r.ClassDecoration),
		RivalCount:          r.RivalCount,
		ReverseRivalCount:   r.ReverseRivalCount,
		MaxStreak:           r.MaxStreak,
		Rank:                r.Rank,
	}
}

func (r userProblemStatsResp) mapToDomainEntity() *domain.UserProblemStats {
	l := make(domain.UserProblemStats, len(r))
	for i, v := range r {
		l[i] = v.mapToDomainEntity()
	}
	return &l
}

func (r userProblemStatResp) mapToDomainEntity() *domain.UserProblemStat {
	return &domain.UserProblemStat{
		Level:   domain.Level(r.Level),
		Total:   r.Total,
		Solved:  r.Solved,
		Partial: r.Partial,
		Tried:   r.Tried,
		Exp:     r.Exp,
	}
}
