package tables

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

var Generators = map[string]table.Generator{
	"persons": GetPersonTable,
	"users":   GetUsersTable,
	"flats":   GetFlatsTable,

	"documents":     GetDocumentsTable,
	"events":        GetEventsTable,
	"events_log":    GetEventsLogTable,
	"faq":           GetFaqTable,
	"faqCategories": GetFaqCategoriesTable,
	"instructions":  GetInstructionsTable,
	"invites":       GetInvitesTable,
	"posts":         GetPostsTable,
	"queue":         GetQueueTable,
	"residents":     GetResidentsTable,
	"roles":         GetRolesTable,
	"sessions":      GetSessionsTable,

	"votes":         GetVotesTable,
	"voteQuestions": GetVoteQuestionsTable,
	"voteAnswers":   GetVoteAnswersTable,
	"votePersons":   GetVotePersonsTable,
}
