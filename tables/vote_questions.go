package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetVoteQuestionsTable(ctx *context.Context) table.Table {

	voteQuestions := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := voteQuestions.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)
	info.AddField("VoteId", "voteId", db.Int4)
	info.AddField("Body", "body", db.Varchar)

	info.SetTable("voteQuestions").SetTitle("VoteQuestions").SetDescription("VoteQuestions")

	formList := voteQuestions.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)
	formList.AddField("VoteId", "voteId", db.Int4, form.Number)
	formList.AddField("Body", "body", db.Varchar, form.Text)

	formList.SetTable("voteQuestions").SetTitle("VoteQuestions").SetDescription("VoteQuestions")

	return voteQuestions
}
