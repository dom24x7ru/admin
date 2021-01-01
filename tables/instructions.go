package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetInstructionsTable(ctx *context.Context) table.Table {

	instructions := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := instructions.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Subtitle", "subtitle", db.Varchar)
	info.AddField("Body", "body", db.JSON)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)

	info.SetTable("instructions").SetTitle("Instructions").SetDescription("Instructions")

	formList := instructions.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Subtitle", "subtitle", db.Varchar, form.Text)
	formList.AddField("Body", "body", db.JSON, form.Text)
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)

	formList.SetTable("instructions").SetTitle("Instructions").SetDescription("Instructions")

	return instructions
}
