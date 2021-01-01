package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetVotesTable(ctx *context.Context) table.Table {

	votes := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := votes.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Multi", "multi", db.Bool)
	info.AddField("Anonymous", "anonymous", db.Bool)
	info.AddField("Closed", "closed", db.Bool)
	info.AddField("House", "house", db.Bool)
	info.AddField("Section", "section", db.Int4)
	info.AddField("Floor", "floor", db.Int4)
	info.AddField("UserId", "userId", db.Int4)

	info.SetTable("votes").SetTitle("Votes").SetDescription("Votes")

	formList := votes.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Multi", "multi", db.Bool, form.Text)
	formList.AddField("Anonymous", "anonymous", db.Bool, form.Text)
	formList.AddField("Closed", "closed", db.Bool, form.Text)
	formList.AddField("House", "house", db.Bool, form.Text)
	formList.AddField("Section", "section", db.Int4, form.Number)
	formList.AddField("Floor", "floor", db.Int4, form.Number)
	formList.AddField("UserId", "userId", db.Int4, form.Number)

	formList.SetTable("votes").SetTitle("Votes").SetDescription("Votes")

	return votes
}
