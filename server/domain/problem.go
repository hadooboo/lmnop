package domain

type Problem struct {
	ProblemID int64
	Level     Level
}

type ProblemsDTO struct {
	Count int64
	Items []*ProblemDTO
}

type ProblemDTO struct {
	ProblemID int64
	Level     Level
}
