package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetQueueTable(ctx *context.Context) table.Table {

	queue := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := queue.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Prefix", "prefix", db.Varchar)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Status", "status", db.Varchar)
	info.AddField("Data", "data", db.Text)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)

	info.SetTable("queue").SetTitle("Queue").SetDescription("Queue")

	formList := queue.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Prefix", "prefix", db.Varchar, form.Text)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Varchar, form.Text)
	formList.AddField("Data", "data", db.Text, form.RichText)
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)

	formList.SetTable("queue").SetTitle("Queue").SetDescription("Queue")

	return queue
}
