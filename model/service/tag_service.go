package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"kry-go/model/entity"
	"kry-go/model/request"
	"kry-go/model/response"
)

type TagService struct {
	DB *gorm.DB
}

func (s *TagService) SelectTagsAt(offset int, limit int) ([]response.TagResponse, error) {
	var tags []entity.Tag
	var tagsResponse []response.TagResponse

	result := s.DB.
		Order("id DESC").
		Offset(offset).
		Limit(limit).
		Find(&tags)

	if result.Error != nil {
		return tagsResponse, result.Error
	}

	for _, e := range tags {
		tagsResponse = append(tagsResponse, response.TagResponse{
			ID:   e.ID,
			Name: e.Name,
		})
	}

	return tagsResponse, nil
}

func (s *TagService) CreateTag(r *request.TagRequest) error {
	tag := entity.Tag{
		Name: r.Name,
	}

	result := s.DB.Create(&tag)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *TagService) SelectTag(id int) (response.TagResponse, error) {
	var tag entity.Tag
	var tagResponse response.TagResponse

	result := s.DB.
		Where("id", id).
		Preload(clause.Associations).
		Find(&tag)

	if result.Error != nil {
		return tagResponse, result.Error
	}

	tagResponse = response.TagResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}

	return tagResponse, nil
}
