package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"kry-go/model/entity"
	"kry-go/model/request"
	"kry-go/model/response"
)

type SubmissionService struct {
	DB *gorm.DB
}

func (s *SubmissionService) SelectSubmissionsByFilter(offset int, id int, language string, limit int) ([]response.SubmissionResponse, error) {
	var submissions []entity.Submission
	var submissionsResponse []response.SubmissionResponse

	clauses := make([]clause.Expression, 0)

	if id != -1 {
		clauses = append(clauses, clause.Eq{Column: "user_id", Value: id})
	}

	if language != "" {
		clauses = append(clauses, clause.Eq{Column: "language", Value: language})
	}

	result := s.DB.
		Clauses(clauses...).
		Order("id").
		Offset(offset).
		Limit(limit).
		Find(&submissions)

	if result.Error != nil {
		return submissionsResponse, result.Error
	}

	for _, e := range submissions {
		submissionsResponse = append(submissionsResponse, response.SubmissionResponse{
			Result:    e.Result,
			Language:  e.Language,
			Code:      e.Code,
			UserID:    e.UserID,
			ProblemID: e.ProblemID,
		})
	}

	return submissionsResponse, nil
}

func (s *SubmissionService) CreateSubmission(r *request.SubmissionRequest) error {
	submission := entity.Submission{
		Result:    r.Result,
		Language:  r.Language,
		Code:      r.Code,
		UserID:    r.UserID,
		ProblemID: r.ProblemID,
	}

	result := s.DB.Omit("User").Omit("Problem").Create(&submission)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *SubmissionService) SelectSubmission(id int) (response.SubmissionResponse, error) {
	var submission entity.Submission
	var submissionResponse response.SubmissionResponse

	result := s.DB.
		Where("id", id).
		Preload(clause.Associations).
		Find(&submission)

	if result.Error != nil {
		return submissionResponse, result.Error
	}

	submissionResponse = response.SubmissionResponse{
		Result:    submission.Result,
		Language:  submission.Language,
		Code:      submission.Code,
		UserID:    submission.UserID,
		ProblemID: submission.ProblemID,
	}

	return submissionResponse, nil
}
