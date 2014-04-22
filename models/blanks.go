package models

func (c *ModelManager) Blanks() *BlankModel {
	blackModel := new(BlankModel)
	blackModel.TableName = "Blanks"
	blackModel.Caption = "Анкеты"
	blackModel.Columns = []string{"id", "name", "contest_id", "columns", "colnames", "types"}
	blackModel.ColNames = []string{"Id", "Название", "Мероприятие", "Колонки", "Названия_колонок", "Типы"}
	tmp := map[string]*Field{
		"id":         {"id", "Id", "serial", false},
		"name":       {"name", "Название", "varchar(128)", false},
		"contest_id": {"contest_id", "Мероприятие", "int", true},
		"columns":    {"columns", "Колонки", "varchar(32)[]", false},
		"colnames":   {"columns", "Названия колонок", "varchar(32)[]", false},
		"types":      {"types", "Типы", "varchar(128)[]", false},
	}
	blackModel.Fields = tmp
	return blackModel
}

type BlankModel struct {
	Entity
}
