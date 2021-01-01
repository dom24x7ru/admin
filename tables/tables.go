package tables

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

var Generators = map[string]table.Generator{
	"persons":        GetPersonTable,
	"users":          GetUsersTable,
	"flats":          GetFlatsTable,
	"invites":        GetInvitesTable,
	"roles":          GetRolesTable,
	"sessions":       GetSessionsTable,
	"documents":      GetDocumentsTable,
	"events":         GetEventsTable,
	"events_log":     GetEventsLogTable,
	"faq":            GetFaqTable,
	"faq_categories": GetFaqCategoriesTable,
	"instructions":   GetInstructionsTable,
	"posts":          GetPostsTable,
	"queue":          GetQueueTable,
	"residents":      GetResidentsTable,
	"votes":          GetVotesTable,
	"vote_questions": GetVoteQuestionsTable,
	"vote_answers":   GetVoteAnswersTable,
	"vote_persons":   GetVotePersonsTable,
}
