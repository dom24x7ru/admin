package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
)

// GetPersonTable return the model of table user.
func GetPersonTable(ctx *context.Context) (personTable table.Table) {
	personTable = table.NewDefaultTable(table.Config{
		Driver:     db.DriverPostgresql,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	info := personTable.GetInfo().SetFilterFormLayout(form.LayoutThreeCol)

	info.SetTable("persons").SetTitle("Люди").SetDescription("Соседи")

	info.AddSelectBox("Пол", types.FieldOptions{
		{Value: "муж", Text: "мужчина"},
		{Value: "жен", Text: "женщина"},
	}, action.FieldFilter("пол"))

	// Список колонок
	info.AddField("ID", "id", db.Int).
		FieldSortable()
	info.AddField("UserID", "userId", db.Int).FieldSortable()
	info.AddField("Имя", "name", db.Varchar).
		FieldEditAble(editType.Text).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Фамилия", "surname", db.Varchar).
		FieldEditAble(editType.Text).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Отчество", "midname", db.Varchar).
		FieldEditAble(editType.Text).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Пол", "sex", db.Varchar).
		FieldEditAble(editType.Text).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Биография", "biography", db.Text).
		FieldEditAble(editType.Text).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Telegram", "telegram", db.Varchar).
		FieldEditAble(editType.Text).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Создано", "createdAt", db.Timestamptz).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange})
	info.AddField("Обновлено", "updatedAt", db.Timestamptz).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange})

	// форма
	formList := personTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("UserID", "userId", db.Int, form.Default)
	formList.AddField("Имя", "name", db.Varchar, form.Text)
	formList.AddField("Фамилия", "surname", db.Varchar, form.Text)
	formList.AddField("Отчество", "midname", db.Varchar, form.Text)
	formList.AddField("Пол", "sex", db.Varchar, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "мужчина", Value: "муж"},
			{Text: "женщина", Value: "жен"},
		}).FieldDefault("муж")
	formList.AddField("Биография", "biography", db.Text, form.TextArea)
	formList.AddField("Telegram", "telegram", db.Varchar, form.Text)
	formList.AddField("Создано", "createdAt", db.Timestamp, form.Default).FieldNotAllowAdd()
	formList.AddField("Обновлено", "updatedAt", db.Timestamp, form.Default).FieldNotAllowAdd()

	personTable.GetForm().SetTabGroups(types.
		NewTabGroups("id", "userId", "createdAt", "updatedAt").
		AddGroup("name", "surname", "midname", "sex", "biography", "telegram")).
		SetTabHeaders("Техническая информация", "Личное")

	formList.SetTable("persons").SetTitle("Люди").SetDescription("Люди")

	//formList.SetPostHook(func(values form2.Values) error {
	//	return nil
	//})

	return
}
