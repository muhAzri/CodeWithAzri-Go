package admin_datamodel

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetCoursesTable(ctx *context.Context) table.Table {

	courses := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("postgresql", "admin_datamodel"))

	info := courses.GetInfo().HideFilterArea()

	info.SetTable("public.courses").SetTitle("Courses").SetDescription("Courses")

	formList := courses.GetForm()

	formList.SetTable("public.courses").SetTitle("Courses").SetDescription("Courses")

	return courses
}
