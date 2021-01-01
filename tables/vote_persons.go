package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetVotePersonsTable(ctx *context.Context) table.Table {

	votePersons := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := votePersons.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)
	info.AddField("VoteId", "voteId", db.Int4)
	info.AddField("PersonId", "personId", db.Int4)

	info.SetTable("votePersons").SetTitle("VotePersons").SetDescription("VotePersons")

	formList := votePersons.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)
	formList.AddField("VoteId", "voteId", db.Int4, form.Number)
	formList.AddField("PersonId", "personId", db.Int4, form.Number)

	formList.SetTable("votePersons").SetTitle("VotePersons").SetDescription("VotePersons")

	return votePersons
}
