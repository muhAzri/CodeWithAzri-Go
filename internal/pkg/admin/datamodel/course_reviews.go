package admin_datamodel

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetCourseReviewsTable(ctx *context.Context) table.Table {

	courseReviews := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("postgresql", "admin_datamodel"))

	info := courseReviews.GetInfo().HideFilterArea()

	info.SetTable("public.course_reviews").SetTitle("CourseReviews").SetDescription("CourseReviews")

	formList := courseReviews.GetForm()

	formList.SetTable("public.course_reviews").SetTitle("CourseReviews").SetDescription("CourseReviews")

	return courseReviews
}
