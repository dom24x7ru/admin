package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetEventsLogTable(ctx *context.Context) table.Table {

	eventsLog := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := eventsLog.GetInfo()

	info.AddField("Id", "id", db.Int4)
	info.AddField("UserId", "userId", db.Int4)
	info.AddField("EventId", "eventId", db.Int4)
	info.AddField("Создано", "Создано", db.Timestamptz)
	info.AddField("Обновлено", "Обновлено", db.Timestamptz)

	info.SetTable("events_log").SetTitle("EventsLog").SetDescription("EventsLog")

	formList := eventsLog.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("UserId", "userId", db.Int4, form.Number)
	formList.AddField("EventId", "eventId", db.Int4, form.Number)
	formList.AddField("Создано", "Создано", db.Timestamptz, form.Datetime)
	formList.AddField("Обновлено", "Обновлено", db.Timestamptz, form.Datetime)

	formList.SetTable("events_log").SetTitle("EventsLog").SetDescription("EventsLog")

	return eventsLog
}
