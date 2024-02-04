package admin_datamodel

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetCourseReviewsCoursesTable(ctx *context.Context) table.Table {

	courseReviewsCourses := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("postgresql", "admin_datamodel"))

	info := courseReviewsCourses.GetInfo().HideFilterArea()

	info.SetTable("public.course_reviews_courses").SetTitle("CourseReviewsCourses").SetDescription("CourseReviewsCourses")

	formList := courseReviewsCourses.GetForm()

	formList.SetTable("public.course_reviews_courses").SetTitle("CourseReviewsCourses").SetDescription("CourseReviewsCourses")

	return courseReviewsCourses
}
