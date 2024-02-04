package admin_datamodel

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetUsersTable(ctx *context.Context) table.Table {

	users := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("postgresql", "admin_datamodel"))

	info := users.GetInfo().HideFilterArea()

	info.SetTable("public.users").SetTitle("Users").SetDescription("Users")

	formList := users.GetForm()

	formList.SetTable("public.users").SetTitle("Users").SetDescription("Users")

	return users
}
