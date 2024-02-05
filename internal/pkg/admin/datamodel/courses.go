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

func GetCoursesTable(ctx *context.Context) table.Table {

	courses := table.NewDefaultTable(table.Config{
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

	info := courses.GetInfo()

	info.SetTable("courses").SetTitle("Courses").SetDescription("Courses")

	info.AddField("ID", "id", db.UUID)
	info.AddField("Name", "name", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldSortable()
	info.AddField("Description", "description", db.Text)
	info.AddField("Language", "language", db.Varchar)
	info.AddField("CreatedAt", "created_at", db.Int).
		FieldSortable()

	info.AddField("UpdatedAt", "updated_at", db.Int).
		FieldSortable()

	timeNow := strconv.FormatInt(timepkg.NowUnixMilli(), 10)

	formList := courses.GetForm()

	formList.SetPreProcessFn(func(values form2.Values) form2.Values {
		values.Add("updated_at", timeNow)
		return values
	})

	formList.AddField("ID", "id", db.UUID, form.Default).FieldNotAllowEdit().FieldDefault(uuid.New().String())
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Text, form.TextArea)
	formList.AddField("Language", "language", db.Varchar, form.Select).
		FieldOptions(types.FieldOptions{
			{Value: "id", Text: "Indonesian"},
			{Value: "en", Text: "English"},
		})
	formList.AddField("CreatedAt", "created_at", db.Int, form.Text).FieldDisplayButCanNotEditWhenCreate().FieldDisplayButCanNotEditWhenUpdate().FieldDefault(timeNow)
	formList.AddField("UpdatedAt", "updated_at", db.Int, form.Text).FieldDisplayButCanNotEditWhenCreate().FieldDisplayButCanNotEditWhenUpdate().FieldDefault(timeNow)

	formList.SetTable("courses").SetTitle("Courses").SetDescription("Courses")

	return courses
}
