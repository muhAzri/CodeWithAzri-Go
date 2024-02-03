package repository_test

import (
	"CodeWithAzri/internal/app/module/course/entity"
	"CodeWithAzri/internal/app/module/course/repository"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

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

	// Test Read One Success
	testReadOneSuccess(t, mock, repo, courseEntity)

	//Test Read One Scan Failure
	testReadOneScanError(t, mock, repo, courseEntity)

	//Test Read One Query Error
	testReadOneQuerryError(t, mock, repo, courseEntity)
}

func testReadOneSuccess(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository, courseEntity entity.Course) {

	// Mocking the database query
	mock.ExpectQuery("SELECT c.id AS course_id, c.name, c.description, c.language, c.created_at, c.updated_at, t.id AS tag_id, t.name AS tag_name, t.created_at, t.updated_at, g.id AS gallery_id, g.url AS gallery_url, g.course_id AS gallery_course_id, g.created_at, g.updated_at, s.id AS section_id, s.name AS section_name, s.course_id AS section_course_id, s.created_at, s.updated_at, l.id AS lesson_id, l.title AS lesson_title, l.video_url AS lesson_video_url, l.course_id AS lesson_course_id, l.course_section_id AS lesson_section_id, l.created_at, l.updated_at FROM courses c LEFT JOIN course_tags_courses tc ON c.id = tc.course_id LEFT JOIN course_tags t ON tc.course_tags_id = t.id LEFT JOIN course_galleries g ON c.id = g.course_id LEFT JOIN course_sections s ON c.id = s.course_id LEFT JOIN course_lessons l ON s.id = l.course_section_id WHERE c.id = $1").
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

func testReadOneScanError(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository, courseEntity entity.Course) {

	rows := sqlmock.NewRows([]string{
		"course_id", "name", "description", "language", "created_at", "updated_at",
		"tag_id", "tag_name", "tag_created_at", "tag_updated_at",
		"gallery_id", "gallery_url", "gallery_course_id", "gallery_created_at", "gallery_updated_at",
		"section_id", "section_name", "section_course_id", "section_created_at", "section_updated_at",
		"lesson_id", "lesson_title", "lesson_video_url", "lesson_course_id", "lesson_section_id", "lesson_created_at", "lesson_updated_at",
	}).AddRow(
		courseEntity.ID, courseEntity.Name, courseEntity.Description, courseEntity.Language, 121212, 121212,
		"invalid id", "", 121212, 121212,
		uuid.Nil, "", uuid.Nil, 0, 0,
		uuid.Nil, "", uuid.Nil, 0, 0,
		uuid.Nil, "", "", uuid.Nil, uuid.Nil, 0, 0,
	)

	mock.ExpectQuery("SELECT c.id AS course_id, c.name, c.description, c.language, c.created_at, c.updated_at, t.id AS tag_id, t.name AS tag_name, t.created_at, t.updated_at, g.id AS gallery_id, g.url AS gallery_url, g.course_id AS gallery_course_id, g.created_at, g.updated_at, s.id AS section_id, s.name AS section_name, s.course_id AS section_course_id, s.created_at, s.updated_at, l.id AS lesson_id, l.title AS lesson_title, l.video_url AS lesson_video_url, l.course_id AS lesson_course_id, l.course_section_id AS lesson_section_id, l.created_at, l.updated_at FROM courses c LEFT JOIN course_tags_courses tc ON c.id = tc.course_id LEFT JOIN course_tags t ON tc.course_tags_id = t.id LEFT JOIN course_galleries g ON c.id = g.course_id LEFT JOIN course_sections s ON c.id = s.course_id LEFT JOIN course_lessons l ON s.id = l.course_section_id WHERE c.id = $1").
		WithArgs(courseEntity.ID).
		WillReturnRows(rows)

	_, err := repo.ReadOne(courseEntity.ID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sql: Scan error on column index 6, name \"tag_id\"")

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("Expectations not met: %v", err)
	}
}

func testReadOneQuerryError(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository, courseEntity entity.Course) {

	mock.ExpectQuery("SELECT c.id AS course_id, c.name, c.description, c.language, c.created_at, c.updated_at, t.id AS tag_id, t.name AS tag_name, t.created_at, t.updated_at, g.id AS gallery_id, g.url AS gallery_url, g.course_id AS gallery_course_id, g.created_at, g.updated_at, s.id AS section_id, s.name AS section_name, s.course_id AS section_course_id, s.created_at, s.updated_at, l.id AS lesson_id, l.title AS lesson_title, l.video_url AS lesson_video_url, l.course_id AS lesson_course_id, l.course_section_id AS lesson_section_id, l.created_at, l.updated_at FROM courses c LEFT JOIN course_tags_courses tc ON c.id = tc.course_id LEFT JOIN course_tags t ON tc.course_tags_id = t.id LEFT JOIN course_galleries g ON c.id = g.course_id LEFT JOIN course_sections s ON c.id = s.course_id LEFT JOIN course_lessons l ON s.id = l.course_section_id WHERE c.id = $1").
		WithArgs(courseEntity.ID).
		WillReturnError(fmt.Errorf("Querry Error"))

	_, err := repo.ReadOne(courseEntity.ID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Querry Error")

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("Expectations not met: %v", err)
	}
}

func TestRepository_ReadMany(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	courseEntity := MockArrayEntity

	// Mocking the database query
	mock.ExpectQuery("SELECT c.id AS course_id, c.name, c.description, c.language, c.created_at, c.updated_at, t.id AS tag_id, t.name AS tag_name, t.created_at, t.updated_at, g.id AS gallery_id, g.url AS gallery_url, g.course_id AS gallery_course_id, g.created_at, g.updated_at FROM courses c LEFT JOIN course_tags_courses tc ON c.id = tc.course_id LEFT JOIN course_tags t ON tc.course_tags_id = t.id LEFT JOIN course_galleries g ON c.id = g.course_id LIMIT $1 OFFSET $2").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(prepareManyRows(courseEntity))

	// Calling the ReadOne method
	result, err := repo.ReadMany(10, 0)
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

func TestRepository_Error(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	// Mocking the database query error
	testReadManyErrorQeury(t, mock, repo)

	// Mocking the database scan error
	testReadManyScanError(t, mock, repo)
}

func testReadManyErrorQeury(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository) {
	mock.ExpectQuery("SELECT c.id AS course_id, c.name, c.description, c.language, c.created_at, c.updated_at, t.id AS tag_id, t.name AS tag_name, t.created_at, t.updated_at, g.id AS gallery_id, g.url AS gallery_url, g.course_id AS gallery_course_id, g.created_at, g.updated_at FROM courses c LEFT JOIN course_tags_courses tc ON c.id = tc.course_id LEFT JOIN course_tags t ON tc.course_tags_id = t.id LEFT JOIN course_galleries g ON c.id = g.course_id LIMIT $1 OFFSET $2").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(fmt.Errorf("Querry Error"))

	_, err := repo.ReadMany(10, 0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Querry Error")

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("Expectations not met: %v", err)
	}
}

func testReadManyScanError(t *testing.T, mock sqlmock.Sqlmock, repo repository.CourseRepository) {
	rows := sqlmock.NewRows([]string{
		"course_id", "name", "description", "language", "created_at", "updated_at",
		"tag_id", "tag_name", "tag_created_at", "tag_updated_at",
		"gallery_id", "gallery_url", "gallery_course_id", "gallery_created_at", "gallery_updated_at",
	}).AddRow(
		"18a95d2f-a941-4a64-bbe5-256be7626db2", "mock Name", "mock desc", "en", 121212, 121212,
		"invalid id", "", 121212, 121212,
		uuid.Nil, "", uuid.Nil, 0, 0,
	)

	mock.ExpectQuery("SELECT c.id AS course_id, c.name, c.description, c.language, c.created_at, c.updated_at, t.id AS tag_id, t.name AS tag_name, t.created_at, t.updated_at, g.id AS gallery_id, g.url AS gallery_url, g.course_id AS gallery_course_id, g.created_at, g.updated_at FROM courses c LEFT JOIN course_tags_courses tc ON c.id = tc.course_id LEFT JOIN course_tags t ON tc.course_tags_id = t.id LEFT JOIN course_galleries g ON c.id = g.course_id LIMIT $1 OFFSET $2").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(rows)

	_, err := repo.ReadMany(10, 0)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Scan error")

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("Expectations not met: %v", err)
	}
}

func TestRepository_Update(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	courseEntity := MockEntity

	mock.ExpectBegin()

	// Expect the course details update query
	mock.ExpectExec(`UPDATE courses SET name = $1, description = $2 , language = $3, updated_at = $4 WHERE id = $5`).WithArgs(
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
	).WillReturnResult(sqlmock.NewResult(0, 1))

	// Expect the deletion of existing tags query
	mock.ExpectExec(`
		DELETE FROM course_tags_courses
		WHERE course_id = $1
	`).WithArgs(sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 1))

	// Expect linking course to tags queries
	for _, tag := range courseEntity.CourseTags {
		mock.ExpectExec(`
		INSERT INTO course_tags_courses (course_id, course_tags_id) VALUES ($1, $2)
		`).WithArgs(courseEntity.ID, tag.ID).WillReturnResult(sqlmock.NewResult(0, 1))
	}

	// Expect gallery update/insert queries
	for _, galleryItem := range courseEntity.Gallery {
		mock.ExpectExec(`
			INSERT INTO course_galleries (id, course_id, url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (id) DO UPDATE SET url = $3, updated_at = $5
		`).WithArgs(
			galleryItem.ID,
			courseEntity.ID,
			galleryItem.URL,
			galleryItem.CreatedAt,
			galleryItem.UpdatedAt,
		).WillReturnResult(sqlmock.NewResult(0, 1))
	}

	// Expect section update/insert queries
	for _, section := range courseEntity.Sections {
		mock.ExpectExec(`
			INSERT INTO course_sections (id, course_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (id) DO UPDATE SET name = $3, updated_at = $5
		`).WithArgs(
			section.ID,
			courseEntity.ID,
			section.Name,
			section.CreatedAt,
			section.UpdatedAt,
		).WillReturnResult(sqlmock.NewResult(0, 1))

		// Expect lesson update/insert queries within each section
		for _, lesson := range section.Lessons {
			mock.ExpectExec(`
			INSERT INTO course_lessons (id, course_id, course_section_id, title, video_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (id) DO UPDATE SET title = $4, video_url = $5, updated_at = $7
			`).WithArgs(
				lesson.ID,
				courseEntity.ID,
				section.ID,
				lesson.Title,
				lesson.VideoURL,
				lesson.CreatedAt,
				lesson.UpdatedAt,
			).WillReturnResult(sqlmock.NewResult(0, 1))
		}
	}

	mock.ExpectCommit()

	// Call the method you are testing
	err := repo.Update(courseEntity.ID, courseEntity)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}

}
