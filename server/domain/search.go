package domain

type Problems struct {
	Count int64
	Items []*Problem
}

type Problem struct {
	ProblemID int64
	Level     Level
}
