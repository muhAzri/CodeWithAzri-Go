package repository

import (
	"CodeWithAzri/internal/app/module/course/entity"
	"database/sql"
	"fmt"

	language_enum "CodeWithAzri/pkg/enums/language"

	"github.com/google/uuid"
)

type CourseRepository interface {
	Create(e entity.Course) error
	ReadMany(limit, offset int) ([]entity.Course, error)
	ReadOne(id uuid.UUID) (entity.Course, error)
	Update(id uuid.UUID, e entity.Course) error
	Delete(id uuid.UUID) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	r := &Repository{db: db}
	return r
}

func (r *Repository) Create(course entity.Course) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	courseQuery := `
		INSERT INTO courses (id, name, description, language, created_at, updated_at)  
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err = tx.Exec(courseQuery, course.ID, course.Name, course.Description, course.Language, course.CreatedAt, course.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create course: %v", err)
	}

	for _, tag := range course.CourseTags {
		linkQuery := `
			INSERT INTO course_tags_courses (course_id, course_tags_id) 
			VALUES ($1, $2)
		`
		_, err = tx.Exec(linkQuery, course.ID, tag.ID)
		if err != nil {
			return fmt.Errorf("failed to link course to tag: %v", err)
		}
	}

	for _, galleryItem := range course.Gallery {
		galleryQuery := `
			INSERT INTO course_galleries (id, course_id, url, created_at, updated_at)  
			VALUES ($1, $2, $3, $4, $5)
		`
		_, err = tx.Exec(galleryQuery, galleryItem.ID, course.ID, galleryItem.URL, galleryItem.CreatedAt, galleryItem.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to create gallery item: %v", err)
		}
	}

	for _, section := range course.Sections {
		sectionQuery := `
			INSERT INTO course_sections (id, course_id, name, created_at, updated_at)   
			VALUES ($1, $2, $3, $4, $5)
		`
		_, err = tx.Exec(sectionQuery, section.ID, course.ID, section.Name, section.CreatedAt, section.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to create section: %v", err)
		}

		for _, lesson := range section.Lessons {
			lessonQuery := `
				INSERT INTO course_lessons (id, course_id, course_section_id, title, video_url, created_at, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)
			`
			_, err = tx.Exec(lessonQuery, lesson.ID, course.ID, section.ID, lesson.Title, lesson.VideoURL, lesson.CreatedAt, lesson.UpdatedAt)
			if err != nil {
				return fmt.Errorf("failed to create lesson: %v", err)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (r *Repository) ReadOne(id uuid.UUID) (entity.Course, error) {
	courseQuery := `
    SELECT c.id AS course_id, c.name, c.description, c.language,
           t.id AS tag_id, t.name AS tag_name,
           g.id AS gallery_id, g.url AS gallery_url, g.course_id AS gallery_course_id,
           s.id AS section_id, s.name AS section_name, s.course_id AS section_course_id,
           l.id AS lesson_id, l.title AS lesson_title, l.video_url AS lesson_video_url, l.course_id AS lesson_course_id, l.course_section_id AS lesson_section_id
    FROM courses c
		LEFT JOIN course_tags_courses tc ON c.id = tc.course_id
		LEFT JOIN course_tags t ON tc.course_tags_id = t.id
		LEFT JOIN course_galleries g ON c.id = g.course_id
		LEFT JOIN course_sections s ON c.id = s.course_id
		LEFT JOIN course_lessons l ON s.id = l.course_section_id
    WHERE c.id = $1
`

	rows, err := r.db.Query(courseQuery, id)
	if err != nil {
		return entity.Course{}, fmt.Errorf("failed to read course: %v", err)
	}
	defer rows.Close()

	course, err := scanReadOne(rows)
	if err != nil {
		return entity.Course{}, fmt.Errorf("failed to scan course: %v", err)
	}

	return course, nil
}

func (r *Repository) ReadMany(limit, offset int) ([]entity.Course, error) {
	var coursesMap = make(map[uuid.UUID]*entity.Course)

	tagExists := func(tags []entity.CourseTags, tagID uuid.UUID) bool {
		for _, tag := range tags {
			if tag.ID == tagID {
				return true
			}
		}
		return false
	}

	galleryExists := func(galleries []entity.CourseGallery, galleryID uuid.UUID) bool {
		for _, gallery := range galleries {
			if gallery.ID == galleryID {
				return true
			}
		}
		return false
	}

	coursesQuery := `
		SELECT c.id AS course_id, c.name, c.description, c.language,
			t.id AS tag_id, t.name AS tag_name,
			g.id AS gallery_id, g.url AS gallery_url, g.course_id AS gallery_course_id
		FROM courses c
			LEFT JOIN course_tags_courses tc ON c.id = tc.course_id
			LEFT JOIN course_tags t ON tc.course_tags_id = t.id
			LEFT JOIN course_galleries g ON c.id = g.course_id
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(coursesQuery, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to read courses: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var courseID, tagID, galleryID, galleryCourseID uuid.UUID
		var courseName, courseDescription, courseLanguage, tagName, galleryURL sql.NullString

		err := rows.Scan(&courseID, &courseName, &courseDescription, &courseLanguage, &tagID, &tagName, &galleryID, &galleryURL, &galleryCourseID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if _, ok := coursesMap[courseID]; !ok {
			coursesMap[courseID] = &entity.Course{
				ID:          courseID,
				Name:        courseName.String,
				Description: courseDescription.String,
				Language:    language_enum.Language(courseLanguage.String),
			}
		}

		if tagID != uuid.Nil && tagName.Valid && !tagExists(coursesMap[courseID].CourseTags, tagID) {
			coursesMap[courseID].CourseTags = append(coursesMap[courseID].CourseTags, entity.CourseTags{
				ID:   tagID,
				Name: tagName.String,
			})
		}

		if galleryID != uuid.Nil && galleryURL.Valid && galleryCourseID != uuid.Nil && !galleryExists(coursesMap[courseID].Gallery, galleryID) {
			coursesMap[courseID].Gallery = append(coursesMap[courseID].Gallery, entity.CourseGallery{
				ID:       galleryID,
				CourseID: galleryCourseID,
				URL:      galleryURL.String,
			})
		}
	}

	var courses []entity.Course
	for _, course := range coursesMap {
		courses = append(courses, *course)
	}

	return courses, nil
}

func (r *Repository) Update(id uuid.UUID, updatedCourse entity.Course) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	courseUpdateQuery := `
		UPDATE courses 
		SET name = $1, description = $2 , language = $3
		WHERE id = $4
	`
	_, err = tx.Exec(courseUpdateQuery, updatedCourse.Name, updatedCourse.Description, updatedCourse.Language, id)
	if err != nil {
		return fmt.Errorf("failed to update course details: %v", err)
	}

	deleteTagsQuery := `
		DELETE FROM course_tags_courses
		WHERE course_id = $1
	`
	_, err = tx.Exec(deleteTagsQuery, id)
	if err != nil {
		return fmt.Errorf("failed to delete existing tags for the course: %v", err)
	}

	for _, tag := range updatedCourse.CourseTags {
		linkQuery := `
			INSERT INTO course_tags_courses (course_id, course_tags_id) 
			VALUES ($1, $2)
		`
		_, err = tx.Exec(linkQuery, id, tag.ID)
		if err != nil {
			return fmt.Errorf("failed to link course to tag: %v", err)
		}
	}

	for _, galleryItem := range updatedCourse.Gallery {
		galleryQuery := `
			INSERT INTO course_galleries (id, course_id, url) 
			VALUES ($1, $2, $3)
			ON CONFLICT (id) DO UPDATE SET url = $3
		`
		_, err = tx.Exec(galleryQuery, galleryItem.ID, id, galleryItem.URL)
		if err != nil {
			return fmt.Errorf("failed to update or insert gallery item: %v", err)
		}
	}

	for _, section := range updatedCourse.Sections {
		sectionQuery := `
			INSERT INTO course_sections (id, course_id, name) 
			VALUES ($1, $2, $3)
			ON CONFLICT (id) DO UPDATE SET name = $3
		`
		_, err = tx.Exec(sectionQuery, section.ID, id, section.Name)
		if err != nil {
			return fmt.Errorf("failed to update or insert section: %v", err)
		}

		for _, lesson := range section.Lessons {
			lessonQuery := `
				INSERT INTO course_lessons (id, course_id, course_section_id, title, video_url) 
				VALUES ($1, $2, $3, $4, $5)
				ON CONFLICT (id) DO UPDATE SET title = $4, video_url = $5
			`
			_, err = tx.Exec(lessonQuery, lesson.ID, id, section.ID, lesson.Title, lesson.VideoURL)
			if err != nil {
				return fmt.Errorf("failed to update or insert lesson: %v", err)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (r *Repository) Delete(id uuid.UUID) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	deleteTagsQuery := `
		DELETE FROM course_tags_courses
		WHERE course_id = $1
	`
	_, err = tx.Exec(deleteTagsQuery, id)
	if err != nil {
		return fmt.Errorf("failed to delete CourseTags associations: %v", err)
	}

	deleteReviewsQuery := `
		DELETE FROM course_reviews_courses
		WHERE course_id = $1
	`
	_, err = tx.Exec(deleteReviewsQuery, id)
	if err != nil {
		return fmt.Errorf("failed to delete CourseReviews associations: %v", err)
	}

	deleteGalleryQuery := `
		DELETE FROM course_galleries
		WHERE course_id = $1
	`
	_, err = tx.Exec(deleteGalleryQuery, id)
	if err != nil {
		return fmt.Errorf("failed to delete CourseGallery: %v", err)
	}

	deleteLessonsSectionsQuery := `
		DELETE FROM course_lessons
		WHERE course_id = $1
	`
	_, err = tx.Exec(deleteLessonsSectionsQuery, id)
	if err != nil {
		return fmt.Errorf("failed to delete CourseLessons: %v", err)
	}

	deleteSectionsQuery := `
		DELETE FROM course_sections
		WHERE course_id = $1
	`
	_, err = tx.Exec(deleteSectionsQuery, id)
	if err != nil {
		return fmt.Errorf("failed to delete CourseSections: %v", err)
	}

	deleteCourseQuery := `
		DELETE FROM courses
		WHERE id = $1
	`
	_, err = tx.Exec(deleteCourseQuery, id)
	if err != nil {
		return fmt.Errorf("failed to delete Course: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func scanReadOne(rows *sql.Rows) (entity.Course, error) {
	var course entity.Course

	var currentCourseID uuid.UUID
	var (
		tagMap     = make(map[uuid.UUID]struct{})
		galleryMap = make(map[uuid.UUID]struct{})
		sectionMap = make(map[uuid.UUID]struct{})
	)

	for rows.Next() {
		var tag entity.CourseTags
		var gallery entity.CourseGallery
		var section entity.CourseSection
		var lesson entity.CourseLesson

		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.Language,
			&tag.ID, &tag.Name,
			&gallery.ID, &gallery.URL, &gallery.CourseID,
			&section.ID, &section.Name, &section.CourseID,
			&lesson.ID, &lesson.Title, &lesson.VideoURL, &lesson.CourseID, &lesson.CourseSectionID,
		)
		if err != nil {
			return entity.Course{}, err
		}

		if currentCourseID != course.ID {
			currentCourseID = course.ID
			course.CourseTags = nil
			course.Gallery = nil
			course.Sections = nil
			tagMap = make(map[uuid.UUID]struct{})
			galleryMap = make(map[uuid.UUID]struct{})
			sectionMap = make(map[uuid.UUID]struct{})
		}

		if tag.ID != uuid.Nil {
			if _, ok := tagMap[tag.ID]; !ok {
				tagMap[tag.ID] = struct{}{}
				course.CourseTags = append(course.CourseTags, tag)
			}
		}

		if gallery.ID != uuid.Nil {
			if _, ok := galleryMap[gallery.ID]; !ok {
				galleryMap[gallery.ID] = struct{}{}
				course.Gallery = append(course.Gallery, gallery)
			}
		}

		if section.ID != uuid.Nil {
			if _, ok := sectionMap[section.ID]; !ok {
				sectionMap[section.ID] = struct{}{}
				currentSection := entity.CourseSection{
					ID:       section.ID,
					Name:     section.Name,
					CourseID: section.CourseID,
					Lessons:  []entity.CourseLesson{lesson},
				}

				course.Sections = append(course.Sections, currentSection)
			} else {
				lessonExists := false
				for i, existingSection := range course.Sections {
					if existingSection.ID == section.ID {
						for _, existingLesson := range existingSection.Lessons {
							if existingLesson.ID == lesson.ID {
								lessonExists = true
								break
							}
						}
						if !lessonExists {
							course.Sections[i].Lessons = append(course.Sections[i].Lessons, lesson)
						}
						break
					}
				}
			}
		}
	}

	return course, nil
}
