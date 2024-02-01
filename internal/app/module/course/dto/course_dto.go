package dto

import (
	language_enum "CodeWithAzri/pkg/enums/language"

	"github.com/google/uuid"
)

type CourseDTO struct {
	ID            uuid.UUID              `json:"id,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Language      language_enum.Language `json:"language,omitempty"`
	CourseTags    []CourseTagsDTO        `json:"tags,omitempty"`
	CourseReviews []CourseReviewsDTO     `json:"reviews,omitempty"`
	Gallery       []CourseGalleryDTO     `json:"gallery,omitempty"`
	Sections      []CourseSectionDTO     `json:"sections,omitempty"`
}

type CourseGalleryDTO struct {
	ID       uuid.UUID `json:"id,omitempty"`
	CourseID uuid.UUID `json:"course_id,omitempty"`
	URL      string    `json:"url,omitempty"`
}

type CourseSectionDTO struct {
	ID       uuid.UUID         `json:"id,omitempty"`
	CourseID uuid.UUID         `json:"course_id,omitempty"`
	Name     string            `json:"name,omitempty"`
	Lessons  []CourseLessonDTO `json:"lessons,omitempty"`
}

type CourseLessonDTO struct {
	ID              uuid.UUID `json:"id,omitempty"`
	CourseID        uuid.UUID `json:"course_id,omitempty"`
	CourseSectionID uuid.UUID `json:"course_section_id,omitempty"`
	Title           string    `json:"title,omitempty"`
	VideoURL        string    `json:"video_url,omitempty"`
}

type CourseReviewsDTO struct {
	ID       uuid.UUID `json:"id,omitempty"`
	CourseID uuid.UUID `json:"course_id,omitempty"`
	UserID   uuid.UUID `json:"user_id,omitempty"`
	Value    int       `json:"value,omitempty"`
	Comment  string    `json:"comment,omitempty"`
}

type CourseTagsDTO struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
}

type CourseIDDTO struct {
	ID uuid.UUID `uri:"id" binding:"required"`
}
