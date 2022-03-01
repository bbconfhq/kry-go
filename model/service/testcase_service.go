package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"kry-go/model/entity"
	"kry-go/model/response"
)

type TestcaseService struct {
	DB *gorm.DB
}

func (s *TestcaseService) SelectTestcasesAt(problemId int) (*[]response.TestcaseResponse, error) {
	var testcases []entity.Testcase
	var testcasesResponse []response.TestcaseResponse

	result := s.DB.
		Where("problem_id", problemId).
		Preload(clause.Associations).
		Find(&testcases)

	if result.Error != nil {
		return &testcasesResponse, result.Error
	}

	for _, e := range testcases {
		testcasesResponse = append(testcasesResponse, response.TestcaseResponse{
			Visible:   e.Visible,
			Input:     e.Input,
			Output:    e.Output,
			ProblemID: e.ProblemID,
		})
	}

	return &testcasesResponse, nil
}
