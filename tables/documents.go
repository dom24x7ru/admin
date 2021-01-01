package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetDocumentsTable(ctx *context.Context) table.Table {

	documents := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := documents.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Заголовок", "title", db.Varchar)
	info.AddField("Аннотация", "annotation", db.Varchar)
	info.AddField("Url", "url", db.Varchar)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)

	info.SetTable("documents").SetTitle("Документы") //.SetDescription("Документы")

	formList := documents.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Заголовок", "title", db.Varchar, form.Text)
	formList.AddField("Аннотация", "annotation", db.Varchar, form.Text)
	formList.AddField("Url", "url", db.Varchar, form.Text)
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)

	formList.SetTable("documents").SetTitle("Документы") //.SetDescription("Документы")

	return documents
}
