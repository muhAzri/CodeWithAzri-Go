package admin_datamodel

import (
	timepkg "CodeWithAzri/pkg/timePkg"
	"strconv"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/google/uuid"
)

func GetCourseLessonsTable(ctx *context.Context) table.Table {

	courseLessons := table.NewDefaultTable(table.Config{
		Driver:     db.DriverPostgresql,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.UUID,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	info := courseLessons.GetInfo()

	info.SetTable("course_lessons").SetTitle("CourseLessons").SetDescription("CourseLessons")

	info.AddField("ID", "id", db.UUID)

	info.AddField("Course Name", "name", db.Varchar).FieldJoin(types.Join{
		Table:     "courses",
		Field:     "course_id",
		JoinField: "id",
	}).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldSortable()

	info.AddField("Course Section Name", "name", db.Varchar).FieldJoin(types.Join{
		Table:     "course_sections",
		Field:     "course_section_id",
		JoinField: "id",
	}).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldSortable()

	info.AddField("Title", "title", db.Varchar)

	info.AddField("CreatedAt", "created_at", db.Int)
	info.AddField("UpdatedAt", "updated_at", db.Int)

	timeNow := strconv.FormatInt(timepkg.NowUnixMilli(), 10)

	formList := courseLessons.GetForm()

	formList.SetPreProcessFn(func(values form2.Values) form2.Values {
		values.Add("updated_at", timeNow)

		return values
	})

	formList.AddField("ID", "id", db.UUID, form.Default).FieldNotAllowEdit().FieldDefault(uuid.New().String())

	formList.AddField("Course", "course_id", db.UUID, form.SelectSingle).FieldOptionsFromTable("courses", "name", "id")
	formList.AddField("Course Section ID", "course_section_id", db.UUID, form.Text)

	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("VideoURL", "video_url", db.Text, form.Text)

	formList.AddField("CreatedAt", "created_at", db.Int, form.Text).FieldDisplayButCanNotEditWhenCreate().FieldDisplayButCanNotEditWhenUpdate().FieldDefault(timeNow)
	formList.AddField("UpdatedAt", "updated_at", db.Int, form.Text).FieldDisplayButCanNotEditWhenCreate().FieldDisplayButCanNotEditWhenUpdate().FieldDefault(timeNow)

	formList.SetTable("course_lessons").SetTitle("CourseLessons").SetDescription("CourseLessons")

	return courseLessons
}
