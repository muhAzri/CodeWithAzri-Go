package admin_datamodel

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
)

func GetCourseTagsCoursesTable(ctx *context.Context) table.Table {

	courseTagsCourses := table.NewDefaultTable(table.Config{
		Driver:     db.DriverPostgresql,
		CanAdd:     true,
		Editable:   false,
		Deletable:  true,
		Exportable: true,
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.UUID,
			Name: "course_id",
		},
	})

	info := courseTagsCourses.GetInfo().HideEditButton()

	info.SetTable("course_tags_courses").SetTitle("CourseTagsCourses").SetDescription("CourseTagsCourses")

	info.AddField("Course ID", "course_id", db.UUID)

	info.AddField("Course Name", "name", db.Varchar).FieldJoin(types.Join{
		Table:     "courses",
		Field:     "course_id",
		JoinField: "id",
	}).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldSortable()

	info.AddField("Tag", "name", db.Varchar).FieldJoin(types.Join{
		Table:     "course_tags",
		Field:     "course_tags_id",
		JoinField: "id",
	}).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldSortable()

	formList := courseTagsCourses.GetForm()

	formList.SetTable("course_tags_courses").SetTitle("CourseTagsCourses").SetDescription("CourseTagsCourses")

	return courseTagsCourses
}
