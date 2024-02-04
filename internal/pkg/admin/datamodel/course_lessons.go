package admin_datamodel

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetCourseLessonsTable(ctx *context.Context) table.Table {

	courseLessons := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("postgresql", "admin_datamodel"))

	info := courseLessons.GetInfo().HideFilterArea()

	info.SetTable("public.course_lessons").SetTitle("CourseLessons").SetDescription("CourseLessons")

	formList := courseLessons.GetForm()

	formList.SetTable("public.course_lessons").SetTitle("CourseLessons").SetDescription("CourseLessons")

	return courseLessons
}
