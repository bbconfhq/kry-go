package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"kry-go/model/entity"
	"kry-go/model/request"
	"kry-go/model/response"
)

type ProblemService struct {
	DB *gorm.DB
}

func (s *ProblemService) SelectProblemsAt(offset int, limit int) (*[]response.ProblemResponse, error) {
	var problems []entity.Problem
	var problemsResponse []response.ProblemResponse

	result := s.DB.
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Select("name")
		}).
		Preload(clause.Associations).
		Order("id").
		Offset(offset).
		Limit(limit).
		Find(&problems)

	if result.Error != nil {
		return &problemsResponse, result.Error
	}

	for _, e := range problems {
		var tags []string

		for _, f := range e.Tags {
			tags = append(tags, f.Name)
		}

		problemsResponse = append(
			problemsResponse,
			response.ProblemResponse{
				ID:          e.ID,
				Title:       e.Title,
				Content:     e.Content,
				Note:        e.Note,
				TimeLimit:   e.TimeLimit,
				MemoryLimit: e.MemoryLimit,
				SubmitCount: e.SubmitCount,
				AcceptCount: e.AcceptCount,
				CreatedAt:   e.CreatedAt,
				UpdatedAt:   e.UpdatedAt,
				Tags:        tags,
			},
		)
	}

	return &problemsResponse, nil
}

func (s *ProblemService) CreateProblem(r *request.ProblemRequest) error {
	var tags []entity.Tag

	for _, e := range r.Tags {
		tags = append(tags, entity.Tag{Name: e})
	}

	problem := entity.Problem{
		Title:       r.Title,
		Content:     r.Content,
		Note:        r.Note,
		TimeLimit:   r.TimeLimit,
		MemoryLimit: r.MemoryLimit,
		Tags:        tags,
	}

	result := s.DB.Omit("Testcases").Create(&problem)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *ProblemService) SelectProblem(id int) (*response.ProblemResponse, error) {
	var problem entity.Problem
	var problemResponse response.ProblemResponse

	result := s.DB.
		Where("id", id).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Select("name")
		}).
		Preload(clause.Associations).
		Order("id").
		Find(&problem)

	if result.Error != nil {
		return &problemResponse, result.Error
	}

	var tags []string
	for _, f := range problem.Tags {
		tags = append(tags, f.Name)
	}

	problemResponse = response.ProblemResponse{
		ID:          problem.ID,
		Title:       problem.Title,
		Content:     problem.Content,
		Note:        problem.Note,
		TimeLimit:   problem.TimeLimit,
		MemoryLimit: problem.MemoryLimit,
		SubmitCount: problem.SubmitCount,
		AcceptCount: problem.AcceptCount,
		CreatedAt:   problem.CreatedAt,
		UpdatedAt:   problem.UpdatedAt,
		Tags:        tags,
	}

	return &problemResponse, nil
}

func (s *ProblemService) UpdateProblem(r *request.ProblemRequest) error {
	var tags []entity.Tag
	for _, e := range r.Tags {
		tags = append(tags, entity.Tag{Name: e})
	}

	problem := entity.Problem{
		Title:       r.Title,
		Content:     r.Content,
		Note:        r.Note,
		TimeLimit:   r.TimeLimit,
		MemoryLimit: r.MemoryLimit,
		Tags:        tags,
	}

	err := s.DB.Model(&entity.Problem{}).Association("Tags").Replace(tags)
	if err != nil {
		return err
	}

	result := s.DB.Model(&entity.Problem{}).Omit("Tags").Updates(problem)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *ProblemService) DeleteProblem(id int) error {
	result := s.DB.Delete(&entity.Problem{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
