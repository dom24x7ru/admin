package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetEventsTable(ctx *context.Context) table.Table {

	events := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := events.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("UserId", "userId", db.Int4)
	info.AddField("Тип", "type", db.Varchar)
	info.AddField("Статус", "status", db.Varchar)
	info.AddField("Приоритет", "priority", db.Varchar)
	info.AddField("Data", "data", db.Text)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)

	info.SetTable("events").SetTitle("События") //.SetDescription("Events")

	formList := events.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("UserId", "userId", db.Int4, form.Number)
	formList.AddField("Тип", "type", db.Varchar, form.Text)
	formList.AddField("Статус", "status", db.Varchar, form.Text)
	formList.AddField("Приоритет", "priority", db.Varchar, form.Text)
	formList.AddField("Data", "data", db.Text, form.RichText)
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)

	formList.SetTable("events").SetTitle("События") //.SetDescription("Events")

	return events
}
