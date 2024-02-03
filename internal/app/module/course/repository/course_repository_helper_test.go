package repository_test

import (
	"CodeWithAzri/internal/app/module/course/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

var mockTags []entity.CourseTags = []entity.CourseTags{
	{
		ID:        uuid.MustParse("345c2c39-5a19-4842-bab8-072a53cd020b"),
		Name:      "Mock Tag",
		CreatedAt: 121212,
		UpdatedAt: 121212,
	},
	{
		ID:        uuid.MustParse("7ccb15a4-483d-4b65-88f8-f2c6d2de3460"),
		Name:      "Mock Tag 2",
		CreatedAt: 121212,
		UpdatedAt: 121212,
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
					ID:              uuid.MustParse("d60619ae-cee9-4877-8f5d-8b294fe9cd81"),
					CourseID:        uuid.MustParse("18a95d2f-a941-4a64-bbe5-256be7626db2"),
					CourseSectionID: uuid.MustParse("b2b71fda-f0f2-4358-9722-b3f13c4564a7"),
					Title:           "Mock Lesson 2",
					VideoURL:        "https://www.youtuber.com",
					CreatedAt:       121212,
					UpdatedAt:       121212,
				},
			},
			CreatedAt: 121212,
			UpdatedAt: 121212,
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
					CreatedAt:       121212,
					UpdatedAt:       121212,
				},
				{
					ID:              uuid.MustParse("d60619ae-cee9-4877-8f5d-8b294fe9cd83"),
					CourseID:        uuid.MustParse("18a95d2f-a941-4a64-bbe5-256be7626db2"),
					CourseSectionID: uuid.MustParse("b2b71fda-f0f2-4358-9722-b3f13c4564a8"),
					Title:           "Mock Lesson 2 2",
					VideoURL:        "https://www.youtuber.com",
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

var MockArrayEntity []entity.Course = []entity.Course{
	{
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
		},
		CreatedAt: 121212,
		UpdatedAt: 121212,
	},
	{
		ID:          uuid.MustParse("a66280a6-61e4-4806-9fc1-8f5457f413a1"),
		Name:        "Mock Course 2 ",
		Description: "Mock Course Description 2",
		Language:    "id",
		CourseTags:  mockTags,
		Gallery: []entity.CourseGallery{
			{
				ID:        uuid.MustParse("d7899f00-3314-487f-a284-75c3916f5605"),
				CourseID:  uuid.MustParse("a66280a6-61e4-4806-9fc1-8f5457f413a1"),
				URL:       "https://www.google.com",
				CreatedAt: 121212,
				UpdatedAt: 121212,
			},
		},
		CreatedAt: 121212,
		UpdatedAt: 121212,
	},
}

func prepareRows(courseEntity entity.Course) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{
		"course_id", "name", "description", "language", "created_at", "updated_at",
		"tag_id", "tag_name", "tag_created_at", "tag_updated_at",
		"gallery_id", "gallery_url", "gallery_course_id", "gallery_created_at", "gallery_updated_at",
		"section_id", "section_name", "section_course_id", "section_created_at", "section_updated_at",
		"lesson_id", "lesson_title", "lesson_video_url", "lesson_course_id", "lesson_section_id", "lesson_created_at", "lesson_updated_at",
	})

	// Adding rows based on the MockEntity
	for _, tag := range courseEntity.CourseTags {
		rows.AddRow(
			courseEntity.ID, courseEntity.Name, courseEntity.Description, courseEntity.Language, 121212, 121212,
			tag.ID, tag.Name, 121212, 121212,
			uuid.Nil, "", uuid.Nil, 0, 0,
			uuid.Nil, "", uuid.Nil, 0, 0,
			uuid.Nil, "", "", uuid.Nil, uuid.Nil, 0, 0,
		)
	}

	for _, gallery := range courseEntity.Gallery {
		rows.AddRow(
			courseEntity.ID, courseEntity.Name, courseEntity.Description, courseEntity.Language, 121212, 121212,
			uuid.Nil, "", 0, 0,
			gallery.ID, gallery.URL, gallery.CourseID, 121212, 121212,
			uuid.Nil, "", uuid.Nil, 0, 0,
			uuid.Nil, "", "", uuid.Nil, uuid.Nil, 0, 0,
		)
	}

	for _, section := range courseEntity.Sections {
		for _, lesson := range section.Lessons {
			rows.AddRow(
				courseEntity.ID, courseEntity.Name, courseEntity.Description, courseEntity.Language, 121212, 121212,
				uuid.Nil, "", 0, 0,
				uuid.Nil, "", uuid.Nil, 0, 0,
				section.ID, section.Name, section.CourseID, 121212, 121212,
				lesson.ID, lesson.Title, lesson.VideoURL, lesson.CourseID, lesson.CourseSectionID, 121212, 121212,
			)
		}
	}

	// //This Additional Row is used to test the case on readOneScan when the lesson is same
	rows.AddRow(
		courseEntity.ID, courseEntity.Name, courseEntity.Description, courseEntity.Language, 121212, 121212,
		uuid.Nil, "", 0, 0,
		uuid.Nil, "", uuid.Nil, 0, 0,
		"b2b71fda-f0f2-4358-9722-b3f13c4564a7", "Mock Section", "18a95d2f-a941-4a64-bbe5-256be7626db2", 121212, 121212,
		"d60619ae-cee9-4877-8f5d-8b294fe9cd80", "Mock Lesson", "https://www.youtube.com", "18a95d2f-a941-4a64-bbe5-256be7626db2", "b2b71fda-f0f2-4358-9722-b3f13c4564a7", 121212, 121212,
	)

	return rows
}

func prepareManyRows(courseArray []entity.Course) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{
		"course_id", "name", "description", "language", "created_at", "updated_at",
		"tag_id", "tag_name", "tag_created_at", "tag_updated_at",
		"gallery_id", "gallery_url", "gallery_course_id", "gallery_created_at", "gallery_updated_at",
	})

	for _, courseEntity := range courseArray {
		for _, tag := range courseEntity.CourseTags {
			rows.AddRow(
				courseEntity.ID, courseEntity.Name, courseEntity.Description, courseEntity.Language, 121212, 121212,
				tag.ID, tag.Name, 121212, 121212,
				uuid.Nil, "", uuid.Nil, 0, 0,
			)
		}

		for _, gallery := range courseEntity.Gallery {
			rows.AddRow(
				courseEntity.ID, courseEntity.Name, courseEntity.Description, courseEntity.Language, 121212, 121212,
				uuid.Nil, "", 0, 0,
				gallery.ID, gallery.URL, gallery.CourseID, 121212, 121212,
			)
		}

	}

	// rows.AddRow(
	// 	courseArray[1].ID, courseArray[1].Name, courseArray[1].Description, courseArray[1].Language, 121212, 121212,
	// 	"345c2c39-5a19-4842-bab8-072a53cd020b", "Mock Tag", 121212, 121212,
	// 	"d7899f00-3314-487f-a284-75c3916f5605", "https://www.google.com", courseArray[1].ID, 121212, 121212,
	// )

	return rows
}
