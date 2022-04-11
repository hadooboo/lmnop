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

func (r userProblemTagStatsResp) mapToDomainEntity() *domain.UserProblemTagStats {
	l := make([]*domain.UserProblemTagStat, len(r.Items))
	for i, v := range r.Items {
		l[i] = v.mapToDomainEntity()
	}
	return &domain.UserProblemTagStats{
		Count: r.Count,
		Items: l,
	}
}

func (r userProblemTagStatResp) mapToDomainEntity() *domain.UserProblemTagStat {
	return &domain.UserProblemTagStat{
		Tag:     r.Tag.Key,
		Total:   r.Total,
		Solved:  r.Solved,
		Partial: r.Partial,
		Tried:   r.Tried,
		Exp:     r.Exp,
	}
}

func (r userTop100Resp) mapToDomainEntity() *domain.UserTop100 {
	l := make([]*domain.Problem, len(r.Items))
	for i, v := range r.Items {
		l[i] = v.mapToDomainEntity()
	}
	return &domain.UserTop100{
		Count: r.Count,
		Items: l,
	}
}

func (r problemResp) mapToDomainEntity() *domain.Problem {
	return &domain.Problem{
		ProblemID: r.ProblemID,
		Level:     domain.Level(r.Level),
	}
}

func (r problemsByTierResp) mapToDomainEntity() *domain.Problems {
	l := make([]*domain.Problem, len(r.Items))
	for i, v := range r.Items {
		l[i] = v.mapToDomainEntity()
	}
	return &domain.Problems{
		Count: r.Count,
		Items: l,
	}
}

func (r problemsByTierAndSolvedByResp) mapToDomainEntity() *domain.Problems {
	l := make([]*domain.Problem, len(r.Items))
	for i, v := range r.Items {
		l[i] = v.mapToDomainEntity()
	}
	return &domain.Problems{
		Count: r.Count,
		Items: l,
	}
}
