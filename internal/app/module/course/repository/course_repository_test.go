package repository_test

import (
	"CodeWithAzri/internal/app/module/course/entity"
	"CodeWithAzri/internal/app/module/course/repository"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var mockTags []entity.CourseTags = []entity.CourseTags{
	{
		ID:        uuid.MustParse("345c2c39-5a19-4842-bab8-072a53cd020b"),
		Name:      "Mock Tag",
		CreatedAt: 0,
		UpdatedAt: 0,
	},
	{
		ID:        uuid.MustParse("7ccb15a4-483d-4b65-88f8-f2c6d2de3460"),
		Name:      "Mock Tag 2",
		CreatedAt: 0,
		UpdatedAt: 0,
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
			CreatedAt: 0,
			UpdatedAt: 0,
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
					CreatedAt:       0,
					UpdatedAt:       0,
				},
				{
					ID:              uuid.MustParse("d60619ae-cee9-4877-8f5d-8b294fe9cd81"),
					CourseID:        uuid.MustParse("18a95d2f-a941-4a64-bbe5-256be7626db2"),
					CourseSectionID: uuid.MustParse("b2b71fda-f0f2-4358-9722-b3f13c4564a7"),
					Title:           "Mock Lesson 2",
					VideoURL:        "https://www.youtuber.com",
					CreatedAt:       0,
					UpdatedAt:       0,
				},
			},
			CreatedAt: 0,
			UpdatedAt: 0,
		},

		{
			ID:       uuid.MustParse("b2b71fda-f0f2-4358-9722-b3f13c4564a8"),
			CourseID: uuid.MustParse("18a95d2f-a941-4a64-bbe5-256be7626db2"),
			Name:     "Mock Section 2",
			Lessons: []entity.CourseLesson{
				{
					ID:              uuid.MustParse("d60619ae-cee9-4877-8f5d-8b294fe9cd82"),
					CourseID:        uuid.MustParse("18a95d2f-a941-4a64-bbe5-256be7626db2"),
					CourseSectionID: uuid.MustParse("b2b71fda-f0f2-4358-9722-b3f13c4564a8"),
					Title:           "Mock Lesson 2 1",
					VideoURL:        "https://www.youtube.com",
					CreatedAt:       0,
					UpdatedAt:       0,
				},
				{
					ID:              uuid.MustParse("d60619ae-cee9-4877-8f5d-8b294fe9cd83"),
					CourseID:        uuid.MustParse("18a95d2f-a941-4a64-bbe5-256be7626db2"),
					CourseSectionID: uuid.MustParse("b2b71fda-f0f2-4358-9722-b3f13c4564a8"),
					Title:           "Mock Lesson 2 2",
					VideoURL:        "https://www.youtuber.com",
					CreatedAt:       0,
					UpdatedAt:       0,
				},
			},
			CreatedAt: 0,
			UpdatedAt: 0,
		},
	},
	CreatedAt: 0,
	UpdatedAt: 0,
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

	//Success
	testCreateSuccess(t, mock, repo, courseEntity)

	//Failed DB Tx Error
	testCreateTransactionErrorHandling(t, mock, repo, courseEntity)

	//Failed DB then Rollback Error
	testCreateRollbackHandling(t, mock, repo, courseEntity)

	//Failed Insert Course
	testCourseInsertErrorHandling(t, mock, repo, courseEntity)

	//Failed to link Course Tag
	testLinkCourseToTagErrorHandling(t, mock, repo, courseEntity)

	//Failed To Create Gallery
	testCreateGalleryItemErrorHandling(t, mock, repo, courseEntity)

	// Test the error handling during creating section
	testCreateSectionErrorHandling(t, mock, repo, courseEntity)

	// Test the error handling during creating lesson
	testCreateLessonErrorHandling(t, mock, repo, courseEntity)

	// Test the error handling during committing the transaction
	testCreateCommitErrorHandling(t, mock, repo, courseEntity)

}

func testCreateSuccess(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository, courseEntity entity.Course) {
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

func testCreateTransactionErrorHandling(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository, courseEntity entity.Course) {
	mock.ExpectBegin().WillReturnError(errors.New("some error"))

	err := repo.Create(courseEntity)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to begin transaction: some error")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func testCreateRollbackHandling(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository, courseEntity entity.Course) {
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

	mock.ExpectExec("INSERT INTO course_tags_courses (course_id, course_tags_id) VALUES ($1, $2)").
		WithArgs(courseEntity.ID, courseEntity.CourseTags[0].ID).
		WillReturnError(errors.New("some error"))

	err := repo.Create(courseEntity)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to link course to tag: some error")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func testCourseInsertErrorHandling(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository, courseEntity entity.Course) {
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
		WillReturnError(errors.New("some error"))

	err := repo.Create(courseEntity)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to create course: some error")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func testLinkCourseToTagErrorHandling(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository, courseEntity entity.Course) {
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

	mock.ExpectExec("INSERT INTO course_tags_courses (course_id, course_tags_id) VALUES ($1, $2)").
		WithArgs(courseEntity.ID, courseEntity.CourseTags[0].ID).
		WillReturnError(errors.New("some error"))

	err := repo.Create(courseEntity)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to link course to tag: some error")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func testCreateGalleryItemErrorHandling(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository, courseEntity entity.Course) {
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

	mock.ExpectExec("INSERT INTO course_galleries (id, course_id, url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)").
		WithArgs(
			courseEntity.Gallery[0].ID,
			courseEntity.ID,
			courseEntity.Gallery[0].URL,
			courseEntity.Gallery[0].CreatedAt,
			courseEntity.Gallery[0].UpdatedAt,
		).
		WillReturnError(errors.New("some error"))

	err := repo.Create(courseEntity)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to create gallery item: some error")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func testCreateSectionErrorHandling(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository, courseEntity entity.Course) {
	// Expectations for the transaction
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

	// Simulate an error during creating section
	mock.ExpectExec("INSERT INTO course_sections (id, course_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)").
		WithArgs(
			courseEntity.Sections[0].ID,
			courseEntity.ID,
			courseEntity.Sections[0].Name,
			courseEntity.Sections[0].CreatedAt,
			courseEntity.Sections[0].UpdatedAt,
		).
		WillReturnError(errors.New("section error"))

	// Call the method being tested
	err := repo.Create(courseEntity)

	// Check if there was an error during the execution
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to create section: section error")

	// Ensure the expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func testCreateLessonErrorHandling(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository, courseEntity entity.Course) {
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

	mock.ExpectExec("INSERT INTO course_sections (id, course_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)").
		WithArgs(
			courseEntity.Sections[0].ID,
			courseEntity.ID,
			courseEntity.Sections[0].Name,
			courseEntity.Sections[0].CreatedAt,
			courseEntity.Sections[0].UpdatedAt,
		).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec("INSERT INTO course_lessons (id, course_id, course_section_id, title, video_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)").
		WithArgs(
			courseEntity.Sections[0].Lessons[0].ID,
			courseEntity.ID,
			courseEntity.Sections[0].ID,
			courseEntity.Sections[0].Lessons[0].Title,
			courseEntity.Sections[0].Lessons[0].VideoURL,
			courseEntity.Sections[0].Lessons[0].CreatedAt,
			courseEntity.Sections[0].Lessons[0].UpdatedAt,
		).
		WillReturnError(errors.New("lesson error"))
	err := repo.Create(courseEntity)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to create lesson: lesson error")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func testCreateCommitErrorHandling(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository, courseEntity entity.Course) {
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

	mock.ExpectCommit().WillReturnError(errors.New("commit error"))

	err := repo.Create(courseEntity)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to commit transaction: commit error")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_ReadOne(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	courseEntity := MockEntity

	testReadOneSuccess(t, mock, repo, courseEntity)
}

func testReadOneSuccess(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository, courseEntity entity.Course) {

	// Mocking the database query
	mock.ExpectQuery("SELECT c.id AS course_id, c.name, c.description, c.language, created_at, updated_at, t.id AS tag_id, t.name AS tag_name, t.created_at, t.updated_at, g.id AS gallery_id, g.url AS gallery_url, g.course_id AS gallery_course_id, g.created_at, g.updated_at, s.id AS section_id, s.name AS section_name, s.course_id AS section_course_id, s.created_at, s.updated_at, l.id AS lesson_id, l.title AS lesson_title, l.video_url AS lesson_video_url, l.course_id AS lesson_course_id, l.course_section_id AS lesson_section_id, l.created_at, l.updated_at FROM courses c LEFT JOIN course_tags_courses tc ON c.id = tc.course_id LEFT JOIN course_tags t ON tc.course_tags_id = t.id LEFT JOIN course_galleries g ON c.id = g.course_id LEFT JOIN course_sections s ON c.id = s.course_id LEFT JOIN course_lessons l ON s.id = l.course_section_id WHERE c.id = $1").
		WithArgs(courseEntity.ID).
		WillReturnRows(prepareRows(courseEntity))

	// Calling the ReadOne method
	result, err := repo.ReadOne(courseEntity.ID)
	if err != nil {
		t.Fatalf("Error while calling ReadOne: %v", err)
	}

	// Asserting the result
	assert.Equal(t, courseEntity, result)

	// Checking if all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("Expectations not met: %v", err)
	}
}

func prepareRows(courseEntity entity.Course) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{
		"course_id", "name", "description", "language",
		"tag_id", "tag_name",
		"gallery_id", "gallery_url", "gallery_course_id",
		"section_id", "section_name", "section_course_id",
		"lesson_id", "lesson_title", "lesson_video_url", "lesson_course_id", "lesson_section_id",
	})

	// Adding rows based on the MockEntity
	for _, tag := range courseEntity.CourseTags {
		rows.AddRow(
			courseEntity.ID, courseEntity.Name, courseEntity.Description, courseEntity.Language,
			tag.ID, tag.Name,
			uuid.Nil, "", uuid.Nil,
			uuid.Nil, "", uuid.Nil,
			uuid.Nil, "", "", uuid.Nil, uuid.Nil,
		)
	}

	for _, gallery := range courseEntity.Gallery {
		rows.AddRow(
			courseEntity.ID, courseEntity.Name, courseEntity.Description, courseEntity.Language,
			uuid.Nil, "",
			gallery.ID, gallery.URL, gallery.CourseID,
			uuid.Nil, "", uuid.Nil,
			uuid.Nil, "", "", uuid.Nil, uuid.Nil,
		)
	}

	for _, section := range courseEntity.Sections {
		for _, lesson := range section.Lessons {
			rows.AddRow(
				courseEntity.ID, courseEntity.Name, courseEntity.Description, courseEntity.Language,
				uuid.Nil, "",
				uuid.Nil, "", uuid.Nil,
				section.ID, section.Name, section.CourseID,
				lesson.ID, lesson.Title, lesson.VideoURL, lesson.CourseID, lesson.CourseSectionID,
			)
		}
	}

	//This Additional Row is used to test the case on readOneScan when the lesson is same
	rows.AddRow(
		courseEntity.ID, courseEntity.Name, courseEntity.Description, courseEntity.Language,
		uuid.Nil, "",
		uuid.Nil, "", uuid.Nil,
		"b2b71fda-f0f2-4358-9722-b3f13c4564a7", "Mock Section", "18a95d2f-a941-4a64-bbe5-256be7626db2",
		"d60619ae-cee9-4877-8f5d-8b294fe9cd80", "Mock Lesson", "https://www.youtube.com", "18a95d2f-a941-4a64-bbe5-256be7626db2", "b2b71fda-f0f2-4358-9722-b3f13c4564a7",
	)

	return rows
}
