package solved_ac_client

import "jaehonam.com/lmnop/domain"

func (r userResp) mapToDomainEntity() *domain.UserDTO {
	return &domain.UserDTO{
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

func (r userProblemStatsResp) mapToDomainEntity() *domain.UserProblemStatsDTO {
	l := make(domain.UserProblemStatsDTO, len(r))
	for i, v := range r {
		l[i] = v.mapToDomainEntity()
	}
	return &l
}

func (r userProblemStatResp) mapToDomainEntity() *domain.UserProblemStatDTO {
	return &domain.UserProblemStatDTO{
		Level:   domain.Level(r.Level),
		Total:   r.Total,
		Solved:  r.Solved,
		Partial: r.Partial,
		Tried:   r.Tried,
		Exp:     r.Exp,
	}
}

func (r userProblemTagStatsResp) mapToDomainEntity() *domain.UserProblemTagStatsDTO {
	l := make([]*domain.UserProblemTagStatDTO, len(r.Items))
	for i, v := range r.Items {
		l[i] = v.mapToDomainEntity()
	}
	return &domain.UserProblemTagStatsDTO{
		Count: r.Count,
		Items: l,
	}
}

func (r userProblemTagStatResp) mapToDomainEntity() *domain.UserProblemTagStatDTO {
	return &domain.UserProblemTagStatDTO{
		Tag:     r.Tag.Key,
		Total:   r.Total,
		Solved:  r.Solved,
		Partial: r.Partial,
		Tried:   r.Tried,
		Exp:     r.Exp,
	}
}

func (r userTop100Resp) mapToDomainEntity() *domain.UserTop100DTO {
	l := make([]*domain.ProblemDTO, len(r.Items))
	for i, v := range r.Items {
		l[i] = v.mapToDomainEntity()
	}
	return &domain.UserTop100DTO{
		Count: r.Count,
		Items: l,
	}
}

func (r problemResp) mapToDomainEntity() *domain.ProblemDTO {
	return &domain.ProblemDTO{
		ProblemID: r.ProblemID,
		Level:     domain.Level(r.Level),
	}
}

func (r problemsByTierResp) mapToDomainEntity() *domain.ProblemsDTO {
	l := make([]*domain.ProblemDTO, len(r.Items))
	for i, v := range r.Items {
		l[i] = v.mapToDomainEntity()
	}
	return &domain.ProblemsDTO{
		Count: r.Count,
		Items: l,
	}
}

func (r problemsByTierAndSolvedByResp) mapToDomainEntity() *domain.ProblemsDTO {
	l := make([]*domain.ProblemDTO, len(r.Items))
	for i, v := range r.Items {
		l[i] = v.mapToDomainEntity()
	}
	return &domain.ProblemsDTO{
		Count: r.Count,
		Items: l,
	}
}
