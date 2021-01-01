package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetVoteAnswersTable(ctx *context.Context) table.Table {

	voteAnswers := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := voteAnswers.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)
	info.AddField("VoteId", "voteId", db.Int4)
	info.AddField("QuestionId", "questionId", db.Int4)
	info.AddField("PersonId", "personId", db.Int4)

	info.SetTable("vote_answers").SetTitle("VoteAnswers").SetDescription("VoteAnswers")

	formList := voteAnswers.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)
	formList.AddField("VoteId", "voteId", db.Int4, form.Number)
	formList.AddField("QuestionId", "questionId", db.Int4, form.Number)
	formList.AddField("PersonId", "personId", db.Int4, form.Number)

	formList.SetTable("vote_answers").SetTitle("VoteAnswers").SetDescription("VoteAnswers")

	return voteAnswers
}
