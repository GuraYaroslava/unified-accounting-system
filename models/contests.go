package models

func (c *ModelManager) Contests() *ContestModel {
    сontestModel := new(ContestModel)
    сontestModel.TableName = "Contests"
    сontestModel.Caption = "Мероприятия"
    сontestModel.Columns = []string{"id", "name", "date"}
    сontestModel.ColNames = []string{"Id", "Название", "Дата"}
    tmp := map[string]*Field{
        "id":   {"id", "Id", "serial", false},
        "name": {"name", "Название", "varchar(128)", false},
        "date": {"date", "Дата создания", "date", false},
    }
    сontestModel.Fields = tmp
    return сontestModel
}

type ContestModel struct {
    Entity
}

type Contest struct {
    Id   string `json:"id"`
    Name string `json:"name"`
    Date string `json:"date"`
}
