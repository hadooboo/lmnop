package application

import (
	"go.uber.org/zap"
	"jaehonam.com/lmnop/application/port/out"
	"jaehonam.com/lmnop/domain"
)

type ProblemService struct {
	client out.ProblemPort
}

func NewProblemService(client out.ProblemPort) *ProblemService {
	return &ProblemService{
		client: client,
	}
}

func (r *ProblemService) GetOptimumProblem(userID string, level domain.Level, except []int64) (*domain.Problem, error) {
	zap.S().Debugf("[service] GetOptimumProblem | userID=%v, level=%v, except=%v", userID, level, except)

	return newOptimumProblemFinder(r.client, userID, level, except).FindOptimumProblem()
}

type optimumProblemFinder struct {
	client                     out.ProblemPort
	userID                     string
	level                      domain.Level
	except                     []int64
	allProblems                []*domain.Problem
	solvedProblems             []*domain.Problem
	allProblemsPointer         int
	solvedProblemsPointer      int
	pageOfAllProblems          int
	isMorePageOfAllProblems    bool
	pageOfSolvedProblems       int
	isMorePageOfSolvedProblems bool
}

func newOptimumProblemFinder(client out.ProblemPort, userID string, level domain.Level, except []int64) *optimumProblemFinder {
	return &optimumProblemFinder{
		client:                     client,
		userID:                     userID,
		level:                      level,
		except:                     except,
		allProblems:                make([]*domain.Problem, 0),
		solvedProblems:             make([]*domain.Problem, 0),
		allProblemsPointer:         0,
		solvedProblemsPointer:      0,
		pageOfAllProblems:          1,
		isMorePageOfAllProblems:    true,
		pageOfSolvedProblems:       1,
		isMorePageOfSolvedProblems: true,
	}
}

func (r *optimumProblemFinder) FindOptimumProblem() (*domain.Problem, error) {
	for {
		if r.isNoMoreAllProblems() && r.isMorePageOfAllProblems {
			if err := r.fillAllProblems(); err != nil {
				return nil, err
			}
		}

		if r.isNoMoreSolvedProblems() && r.isMorePageOfSolvedProblems {
			if err := r.fillSolvedProblems(); err != nil {
				return nil, err
			}
		}

		allProblem, solvedProblem := r.getNextAllProblem(), r.getNextSolvedProblem()
		switch {
		case allProblem == nil:
			return nil, domain.ErrNotFoundProblem
		case solvedProblem != nil && allProblem.ProblemID == solvedProblem.ProblemID:
			r.allProblemsPointer++
			r.solvedProblemsPointer++
		case r.isInExcept(allProblem.ProblemID):
			r.allProblemsPointer++
		default:
			return allProblem, nil
		}
	}
}

func (r *optimumProblemFinder) isNoMoreAllProblems() bool {
	return r.allProblemsPointer == len(r.allProblems)
}

func (r *optimumProblemFinder) fillAllProblems() error {
	problemsByTierDTO, err := r.client.GetProblemsByTier(r.level, r.pageOfAllProblems)
	if err != nil {
		return err
	}
	if len(problemsByTierDTO.Items) == 0 {
		r.isMorePageOfAllProblems = false
		return nil
	}

	for _, v := range problemsByTierDTO.Items {
		r.allProblems = append(r.allProblems, &domain.Problem{
			ProblemID: v.ProblemID,
			Level:     v.Level,
		})
	}

	r.pageOfAllProblems++
	return nil
}

func (r *optimumProblemFinder) isNoMoreSolvedProblems() bool {
	return r.solvedProblemsPointer == len(r.solvedProblems)
}

func (r *optimumProblemFinder) fillSolvedProblems() error {
	problemsByTierAndSolvedByDTO, err := r.client.GetProblemsByTierAndSolvedBy(r.level, r.userID, r.pageOfSolvedProblems)
	if err != nil {
		return err
	}
	if len(problemsByTierAndSolvedByDTO.Items) == 0 {
		r.isMorePageOfSolvedProblems = false
		return nil
	}

	for _, v := range problemsByTierAndSolvedByDTO.Items {
		r.solvedProblems = append(r.solvedProblems, &domain.Problem{
			ProblemID: v.ProblemID,
			Level:     v.Level,
		})
	}

	r.pageOfSolvedProblems++
	return nil
}

func (r *optimumProblemFinder) getNextAllProblem() *domain.Problem {
	if r.isNoMoreAllProblems() {
		return nil
	}
	return r.allProblems[r.allProblemsPointer]
}

func (r *optimumProblemFinder) getNextSolvedProblem() *domain.Problem {
	if r.isNoMoreSolvedProblems() {
		return nil
	}
	return r.solvedProblems[r.solvedProblemsPointer]
}

func (r *optimumProblemFinder) isInExcept(problemID int64) bool {
	for _, v := range r.except {
		if v == problemID {
			return true
		}
	}
	return false
}
