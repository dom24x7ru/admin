package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetFaqTable(ctx *context.Context) table.Table {

	faq := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := faq.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)
	info.AddField("CategoryId", "categoryId", db.Int4)
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Body", "body", db.Text)

	info.SetTable("faq").SetTitle("Faq").SetDescription("Faq")

	formList := faq.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)
	formList.AddField("CategoryId", "categoryId", db.Int4, form.Number)
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Body", "body", db.Text, form.RichText)

	formList.SetTable("faq").SetTitle("Faq").SetDescription("Faq")

	return faq
}
