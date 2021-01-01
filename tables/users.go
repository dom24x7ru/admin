package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUsersTable(ctx *context.Context) table.Table {

	users := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := users.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Mobile", "mobile", db.Varchar)
	info.AddField("SmsCode", "smsCode", db.Varchar)
	info.AddField("Banned", "banned", db.Bool)
	info.AddField("RoleId", "roleId", db.Int4)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)

	info.SetTable("users").SetTitle("Users").SetDescription("Users")

	formList := users.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Mobile", "mobile", db.Varchar, form.Text)
	formList.AddField("SmsCode", "smsCode", db.Varchar, form.Text)
	formList.AddField("Banned", "banned", db.Bool, form.Text)
	formList.AddField("RoleId", "roleId", db.Int4, form.Number)
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)

	formList.SetTable("users").SetTitle("Users").SetDescription("Users")

	return users
}
