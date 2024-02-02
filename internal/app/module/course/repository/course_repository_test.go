package repository_test

import (
	"CodeWithAzri/internal/app/module/course/entity"
	"CodeWithAzri/internal/app/module/course/repository"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var mockTags []entity.CourseTags = []entity.CourseTags{
	{
		ID:   uuid.MustParse("345c2c39-5a19-4842-bab8-072a53cd020b"),
		Name: "Mock Tag",
	},
	{
		ID:   uuid.MustParse("7ccb15a4-483d-4b65-88f8-f2c6d2de3460"),
		Name: "Mock Tag 2",
	},
}

var MockEntity entity.Course = entity.Course{
	ID:          uuid.MustParse("18a95d2f-a941-4a64-bbe5-256be7626db2"),
	Name:        "Mock Course",
	Description: "Mock Course Description",
	Language:    "en",
	CourseTags:  mockTags,
	Gallery: []entity.CourseGallery{
		{
			ID:        uuid.MustParse("b2b71fda-f0f2-4358-9722-b3f13c4564a5"),
			CourseID:  uuid.MustParse("18a95d2f-a941-4a64-bbe5-256be7626db2"),
			URL:       "https://www.google.com",
			CreatedAt: 121212,
			UpdatedAt: 121212,
		},
		{
			ID:        uuid.MustParse("b2b71fda-f0f2-4358-9722-b3f13c4564a6"),
			CourseID:  uuid.MustParse("18a95d2f-a941-4a64-bbe5-256be7626db2"),
			URL:       "https://www.yahoo.com",
			CreatedAt: 121212,
			UpdatedAt: 121212,
		},
	},
	Sections: []entity.CourseSection{
		{
			ID:       uuid.MustParse("b2b71fda-f0f2-4358-9722-b3f13c4564a7"),
			CourseID: uuid.MustParse("18a95d2f-a941-4a64-bbe5-256be7626db2"),
			Name:     "Mock Section",
			Lessons: []entity.CourseLesson{
				{
					ID:              uuid.MustParse("d60619ae-cee9-4877-8f5d-8b294fe9cd80"),
					CourseID:        uuid.MustParse("18a95d2f-a941-4a64-bbe5-256be7626db2"),
					CourseSectionID: uuid.MustParse("b2b71fda-f0f2-4358-9722-b3f13c4564a7"),
					Title:           "Mock Lesson",
					VideoURL:        "https://www.youtube.com",
					CreatedAt:       121212,
					UpdatedAt:       121212,
				},
				{
					ID:              uuid.MustParse("61467da5-d1b9-4fd2-bc53-1f86842ecf77"),
					CourseID:        uuid.MustParse("18a95d2f-a941-4a64-bbe5-256be7626db2"),
					CourseSectionID: uuid.MustParse("b2b71fda-f0f2-4358-9722-b3f13c4564a7"),
					Title:           "Mock Lesson 2",
					VideoURL:        "https://www.youtube.com",
					CreatedAt:       121212,
					UpdatedAt:       121212,
				},
			},
			CreatedAt: 121212,
			UpdatedAt: 121212,
		},
	},
	CreatedAt: 121212,
	UpdatedAt: 121212,
}

func initializeMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repository.CourseRepository) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}

	mockRepo := repository.NewRepository(db)

	return db, mock, mockRepo
}

func TestRepository_Create(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	courseEntity := MockEntity

	mock.ExpectBegin()

	mock.ExpectExec("INSERT INTO courses (id, name, description, language, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)").
		WithArgs(
			courseEntity.ID,
			courseEntity.Name,
			courseEntity.Description,
			courseEntity.Language,
			courseEntity.CreatedAt,
			courseEntity.UpdatedAt,
		).
		WillReturnResult(sqlmock.NewResult(0, 1))

	for _, tag := range courseEntity.CourseTags {
		mock.ExpectExec("INSERT INTO course_tags_courses (course_id, course_tags_id) VALUES ($1, $2)").
			WithArgs(courseEntity.ID, tag.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
	}

	for _, galleryItem := range courseEntity.Gallery {
		mock.ExpectExec("INSERT INTO course_galleries (id, course_id, url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)").
			WithArgs(galleryItem.ID, courseEntity.ID, galleryItem.URL, galleryItem.CreatedAt, galleryItem.UpdatedAt).
			WillReturnResult(sqlmock.NewResult(0, 1))
	}

	for _, section := range courseEntity.Sections {
		mock.ExpectExec("INSERT INTO course_sections (id, course_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)").
			WithArgs(section.ID, courseEntity.ID, section.Name, section.CreatedAt, section.UpdatedAt).
			WillReturnResult(sqlmock.NewResult(0, 1))

		for _, lesson := range section.Lessons {
			mock.ExpectExec("INSERT INTO course_lessons (id, course_id, course_section_id, title, video_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)").
				WithArgs(lesson.ID, courseEntity.ID, section.ID, lesson.Title, lesson.VideoURL, lesson.CreatedAt, lesson.UpdatedAt).
				WillReturnResult(sqlmock.NewResult(0, 1))
		}
	}

	mock.ExpectCommit()

	err := repo.Create(courseEntity)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
