package service

import (
	"CodeWithAzri/internal/app/module/course/dto"
	"CodeWithAzri/internal/app/module/course/repository"
	"CodeWithAzri/pkg/adapter"

	"github.com/google/uuid"
)

type CourseService interface {
	GetDetailCourse(courseID uuid.UUID) (dto.CourseDTO, error)
	GetPaginatedCourses(limit int, page int) ([]dto.CourseDTO, error)
}

type Service struct {
	repository repository.CourseRepository
}

func NewCourseService(r repository.CourseRepository) CourseService {
	s := new(Service)
	s.repository = r
	return s
}

func (s *Service) GetDetailCourse(courseID uuid.UUID) (dto.CourseDTO, error) {
	course, err := s.repository.ReadOne(courseID)
	if err != nil {
		return dto.CourseDTO{}, err
	}

	courseDTO, err := adapter.AnyToType[dto.CourseDTO](course)
	if err != nil {
		return dto.CourseDTO{}, err
	}

	return courseDTO, nil
}

func (s *Service) GetPaginatedCourses(limit int, page int) ([]dto.CourseDTO, error) {
	offset := (page - 1) * limit

	courses, err := s.repository.ReadMany(limit, offset)
	if err != nil {
		return []dto.CourseDTO{}, err
	}

	courseDTOs, err := adapter.AnyToType[[]dto.CourseDTO](courses)
	if err != nil || courseDTOs == nil {
		return []dto.CourseDTO{}, err
	}

	return courseDTOs, nil
}
