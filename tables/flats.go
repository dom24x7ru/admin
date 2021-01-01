package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetFlatsTable(ctx *context.Context) table.Table {

	flats := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := flats.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Number", "number", db.Int4)
	info.AddField("Section", "section", db.Int4)
	info.AddField("Floor", "floor", db.Int4)
	info.AddField("Rooms", "rooms", db.Int4)
	info.AddField("Square", "square", db.Float8)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)

	info.SetTable("flats").SetTitle("Flats").SetDescription("Flats")

	formList := flats.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Number", "number", db.Int4, form.Number)
	formList.AddField("Section", "section", db.Int4, form.Number)
	formList.AddField("Floor", "floor", db.Int4, form.Number)
	formList.AddField("Rooms", "rooms", db.Int4, form.Number)
	formList.AddField("Square", "square", db.Float8, form.Text)
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)

	formList.SetTable("flats").SetTitle("Flats").SetDescription("Flats")

	return flats
}
