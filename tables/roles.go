package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetRolesTable(ctx *context.Context) table.Table {

	roles := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := roles.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Code", "code", db.Varchar)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)

	info.SetTable("roles").SetTitle("Roles").SetDescription("Roles")

	formList := roles.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Code", "code", db.Varchar, form.Text)
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)

	formList.SetTable("roles").SetTitle("Roles").SetDescription("Roles")

	return roles
}
