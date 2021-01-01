package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetResidentsTable(ctx *context.Context) table.Table {

	residents := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := residents.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("PersonId", "personId", db.Int4)
	info.AddField("FlatId", "flatId", db.Int4)
	info.AddField("IsOwner", "isOwner", db.Bool)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)

	info.SetTable("residents").SetTitle("Residents").SetDescription("Residents")

	formList := residents.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("PersonId", "personId", db.Int4, form.Number)
	formList.AddField("FlatId", "flatId", db.Int4, form.Number)
	formList.AddField("IsOwner", "isOwner", db.Bool, form.Text)
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)

	formList.SetTable("residents").SetTitle("Residents").SetDescription("Residents")

	return residents
}
