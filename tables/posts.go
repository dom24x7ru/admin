package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetPostsTable(ctx *context.Context) table.Table {

	posts := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := posts.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Body", "body", db.Text)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)
	info.AddField("Type", "type", db.Varchar)
	info.AddField("Url", "url", db.Varchar)

	info.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	formList := posts.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Body", "body", db.Text, form.RichText)
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)
	formList.AddField("Type", "type", db.Varchar, form.Text)
	formList.AddField("Url", "url", db.Varchar, form.Text)

	formList.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	return posts
}
