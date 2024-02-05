package admin_datamodel

import (
	timepkg "CodeWithAzri/pkg/timePkg"
	"strconv"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/google/uuid"

	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
)

func GetCourseTagsTable(ctx *context.Context) table.Table {

	courseTags := table.NewDefaultTable(table.Config{
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

	info := courseTags.GetInfo().HideFilterArea()

	info.SetTable("course_tags").SetTitle("CourseTags").SetDescription("CourseTags")

	info.AddField("ID", "id", db.UUID)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("CreatedAt", "created_at", db.Int)
	info.AddField("UpdatedAt", "updated_at", db.Int)

	timeNow := strconv.FormatInt(timepkg.NowUnixMilli(), 10)

	formList := courseTags.GetForm()

	formList.SetPreProcessFn(func(values form2.Values) form2.Values {
		values.Add("updated_at", timeNow)
		return values
	})

	formList.AddField("ID", "id", db.UUID, form.Default).FieldNotAllowEdit().FieldDefault(uuid.New().String())
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("CreatedAt", "created_at", db.Int, form.Text).FieldDisplayButCanNotEditWhenCreate().FieldDisplayButCanNotEditWhenUpdate().FieldDefault(timeNow)
	formList.AddField("UpdatedAt", "updated_at", db.Int, form.Text).FieldDisplayButCanNotEditWhenCreate().FieldDisplayButCanNotEditWhenUpdate().FieldDefault(timeNow)

	formList.SetTable("course_tags").SetTitle("CourseTags").SetDescription("CourseTags")

	return courseTags
}
