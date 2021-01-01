package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetInvitesTable(ctx *context.Context) table.Table {

	invites := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := invites.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("UserId", "userId", db.Int4)
	info.AddField("Code", "code", db.Varchar)
	info.AddField("Used", "used", db.Bool)
	info.AddField("NewUserId", "newUserId", db.Int4)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)

	info.SetTable("invites").SetTitle("Invites").SetDescription("Invites")

	formList := invites.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("UserId", "userId", db.Int4, form.Number)
	formList.AddField("Code", "code", db.Varchar, form.Text)
	formList.AddField("Used", "used", db.Bool, form.Text)
	formList.AddField("NewUserId", "newUserId", db.Int4, form.Number)
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)

	formList.SetTable("invites").SetTitle("Invites").SetDescription("Invites")

	return invites
}
