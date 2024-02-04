package admin_datamodel

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetCourseGalleriesTable(ctx *context.Context) table.Table {

	courseGalleries := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("postgresql", "admin_datamodel"))

	info := courseGalleries.GetInfo().HideFilterArea()

	info.SetTable("public.course_galleries").SetTitle("CourseGalleries").SetDescription("CourseGalleries")

	formList := courseGalleries.GetForm()

	formList.SetTable("public.course_galleries").SetTitle("CourseGalleries").SetDescription("CourseGalleries")

	return courseGalleries
}
