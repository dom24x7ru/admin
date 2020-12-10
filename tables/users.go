package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUsersTable(ctx *context.Context) table.Table {

	users := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := users.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Number", "number", db.Int4)
	info.AddField("Section", "section", db.Int4)
	info.AddField("Floor", "floor", db.Int4)
	info.AddField("Rooms", "rooms", db.Int4)
	info.AddField("Square", "square", db.Float8)
	info.AddField("CreatedAt", "createdAt", db.Timestamptz)
	info.AddField("UpdatedAt", "updatedAt", db.Timestamptz)

	info.SetTable("users").SetTitle("Users").SetDescription("Users")

	formList := users.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Number", "number", db.Int4, form.Number)
	formList.AddField("Section", "section", db.Int4, form.Number)
	formList.AddField("Floor", "floor", db.Int4, form.Number)
	formList.AddField("Rooms", "rooms", db.Int4, form.Number)
	formList.AddField("Square", "square", db.Float8, form.Text)
	formList.AddField("CreatedAt", "createdAt", db.Timestamptz, form.Datetime)
	formList.AddField("UpdatedAt", "updatedAt", db.Timestamptz, form.Datetime)

	formList.SetTable("users").SetTitle("Users").SetDescription("Users")

	return users
}
