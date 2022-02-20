package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"kry-go/model/entity"
	"kry-go/model/request"
	"kry-go/model/response"
)

type ContestService struct {
	DB *gorm.DB
}

func (c *ContestService) SelectContestsAt(offset int, limit int) ([]response.ContestResponse, error) {
	var contests []entity.Contest
	var contestsResponse []response.ContestResponse

	c.DB.
		Preload("Problems", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID")
		}).
		Preload(clause.Associations).
		Order("ended_at DESC, started_at DESC, id DESC").
		Offset(offset).
		Limit(limit).
		Find(&contests)

	for _, e := range contests {
		var problemIds []uint

		for _, f := range e.Problems {
			problemIds = append(problemIds, f.ID)
		}

		contestsResponse = append(
			contestsResponse,
			response.ContestResponse{
				ID:         e.ID,
				Title:      e.Title,
				Content:    e.Content,
				ProblemIds: problemIds,
				StartedAt:  e.StartedAt,
				EndedAt:    e.EndedAt,
			},
		)
	}

	return contestsResponse, nil
}

func (c *ContestService) SelectContest(id int) (response.ContestResponse, error) {
	var contest entity.Contest
	var contestResponse response.ContestResponse

	c.DB.
		Where("ID", id).
		Preload("Problems", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID")
		}).
		Preload(clause.Associations).
		Order("ended_at DESC, started_at DESC, id DESC").
		Find(&contest)

	var problemIds []uint
	for _, f := range contest.Problems {
		problemIds = append(problemIds, f.ID)
	}

	contestResponse = response.ContestResponse{
		ID:         contest.ID,
		Title:      contest.Title,
		Content:    contest.Content,
		ProblemIds: problemIds,
		StartedAt:  contest.StartedAt,
		EndedAt:    contest.EndedAt,
	}

	return contestResponse, nil
}

func (c *ContestService) CreateContest(r *request.ContestRequest) error {
	var problem []entity.Problem
	for _, e := range r.ProblemIds {
		problem = append(problem, entity.Problem{ID: e})
	}

	contest := entity.Contest{
		Title:     r.Title,
		Content:   r.Content,
		Problems:  problem,
		StartedAt: r.StartedAt,
		EndedAt:   r.EndedAt,
	}

	c.DB.Omit("Problems.*").Create(&contest)

	return nil
}

func (c *ContestService) UpdateContest(r *request.ContestRequest) error {
	var problem []entity.Problem
	for _, e := range r.ProblemIds {
		problem = append(problem, entity.Problem{ID: e})
	}

	contest := entity.Contest{
		Title:     r.Title,
		Content:   r.Content,
		Problems:  problem,
		StartedAt: r.StartedAt,
		EndedAt:   r.EndedAt,
	}

	c.DB.Model(&entity.Contest{}).Association("Problems").Replace(problem)
	c.DB.Model(&entity.Contest{}).Omit("Problems").Updates(contest)

	return nil
}

func (c *ContestService) DeleteContest(id int) error {
	c.DB.Delete(&entity.Contest{}, id)
	return nil
}
