package service_test

import (
	"CodeWithAzri/internal/app/module/course/entity"

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
		Name:        "Mock Course 2",
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
