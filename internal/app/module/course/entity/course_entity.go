package entity

import (
	language_enum "CodeWithAzri/pkg/enums/language"

	"github.com/google/uuid"
)

type Course struct {
	ID            uuid.UUID              `json:"id" gorm:"type:uuid;primaryKey"`
	Name          string                 `json:"name" gorm:"type:varchar(255)"`
	Description   string                 `json:"description" gorm:"type:text"`
	Language      language_enum.Language `json:"language" gorm:"type:varchar(2)"`
	CourseTags    []CourseTags           `json:"tags" gorm:"many2many:course_tags_courses;"`
	CourseReviews []CourseReviews        `json:"reviews" gorm:"many2many:course_reviews_courses"`
	Gallery       []CourseGallery        `json:"gallery"`
	Sections      []CourseSection        `json:"sections"`
}

type CourseGallery struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	CourseID uuid.UUID `json:"course_id" gorm:"type:uuid;index"`
	URL      string    `json:"url" gorm:"type:text"`
}

type CourseSection struct {
	ID       uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	CourseID uuid.UUID      `json:"course_id" gorm:"type:uuid;index"`
	Name     string         `json:"name" gorm:"type:varchar(255)"`
	Lessons  []CourseLesson `json:"lessons"`
}

type CourseLesson struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	CourseID        uuid.UUID `json:"course_id" gorm:"type:uuid;index"`
	CourseSectionID uuid.UUID `json:"course_section_id" gorm:"type:uuid;index"`
	Title           string    `json:"title" gorm:"type:varchar(255)"`
	VideoURL        string    `json:"video_url" gorm:"type:text"`
}

type CourseReviews struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	CourseID uuid.UUID `json:"course_id" gorm:"type:uuid;index"`
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;index"`
	Value    int       `json:"value"`
	Comment  string    `json:"comment" gorm:"type:text"`
}

type CourseTags struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name string    `json:"name" gorm:"type:varchar(255);uniqueIndex"`
}
