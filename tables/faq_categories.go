package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetFaqCategoriesTable(ctx *context.Context) table.Table {

	faqCategories := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := faqCategories.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)
	info.AddField("Название", "name", db.Varchar)
	info.AddField("Описание", "description", db.Varchar)

	info.SetTable("faq_categories").SetTitle("FAQ категории").SetDescription("FAQ категории")

	formList := faqCategories.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)
	formList.AddField("Название", "name", db.Varchar, form.Text)
	formList.AddField("Описание", "description", db.Varchar, form.Text)

	formList.SetTable("faq_categories").SetTitle("FAQ категории").SetDescription("FAQ категории")

	return faqCategories
}
