package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetSessionsTable(ctx *context.Context) table.Table {

	sessions := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := sessions.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Uuid", "uuid", db.Varchar)
	info.AddField("UserId", "userId", db.Int4)
	info.AddField("Ip", "ip", db.Varchar)
	info.AddField("ForwardedIp", "forwardedIp", db.Varchar)
	info.AddField("Login", "login", db.Timestamptz)
	info.AddField("Logout", "logout", db.Timestamptz)
	info.AddField("Online", "online", db.Bool)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)

	info.SetTable("sessions").SetTitle("Sessions").SetDescription("Sessions")

	formList := sessions.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Uuid", "uuid", db.Varchar, form.Text)
	formList.AddField("UserId", "userId", db.Int4, form.Number)
	formList.AddField("Ip", "ip", db.Varchar, form.Ip)
	formList.AddField("ForwardedIp", "forwardedIp", db.Varchar, form.Text)
	formList.AddField("Login", "login", db.Timestamptz, form.Datetime)
	formList.AddField("Logout", "logout", db.Timestamptz, form.Datetime)
	formList.AddField("Online", "online", db.Bool, form.Text)
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)

	formList.SetTable("sessions").SetTitle("Sessions").SetDescription("Sessions")

	return sessions
}
