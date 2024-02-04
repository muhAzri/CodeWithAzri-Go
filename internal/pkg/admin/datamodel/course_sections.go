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

func GetCourseSectionsTable(ctx *context.Context) table.Table {

	courseSections := table.NewDefaultTable(table.Config{
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

	info := courseSections.GetInfo()

	info.SetTable("course_sections").SetTitle("CourseSections").SetDescription("CourseSections")

	info.AddField("ID", "id", db.UUID)

	info.AddField("Course Name", "name", db.Varchar).FieldJoin(types.Join{
		Table:     "courses",
		Field:     "course_id",
		JoinField: "id",
	}).
	FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).FieldSortable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("CreatedAt", "created_at", db.Int)
	info.AddField("UpdatedAt", "updated_at", db.Int)

	timeNow := strconv.FormatInt(timepkg.NowUnixMilli(), 10)

	formList := courseSections.GetForm()

	formList.SetPreProcessFn(func(values form2.Values) form2.Values {
		values.Add("updated_at", timeNow)
		return values
	})

	formList.AddField("ID", "id", db.UUID, form.Default).FieldNotAllowEdit().FieldDefault(uuid.New().String())

	formList.AddField("Course", "course_id", db.UUID, form.SelectSingle).FieldOptionsFromTable("courses", "name", "id")
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("CreatedAt", "created_at", db.Int, form.Text).FieldDisplayButCanNotEditWhenCreate().FieldDisplayButCanNotEditWhenUpdate().FieldDefault(timeNow)
	formList.AddField("UpdatedAt", "updated_at", db.Int, form.Text).FieldDisplayButCanNotEditWhenCreate().FieldDisplayButCanNotEditWhenUpdate().FieldDefault(timeNow)

	formList.SetTable("course_sections").SetTitle("CourseSections").SetDescription("CourseSections")

	return courseSections
}

// formList.AddField("Course", "course_id", db.UUID, form.SelectSingle).FieldOptionsFromTable("courses", "name", "id", func(sql *db.SQL) *db.SQL {
// 	return sql.Where("name", "LIKE", "%%Test%")
// })
