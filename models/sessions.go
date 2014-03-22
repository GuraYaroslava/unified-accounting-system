package models

func (c *ModelManager) Sessions() *SessionModel {
    sessionModel := new(SessionModel)

    sessionModel.TableName = "Sessions"
    sessionModel.Caption = "Сессии"
    sessionModel.Columns = []string{"sid", "login"}
    tmp := map[string]*Field{
        "sid":   {"sid", "Sid", "integer", false},
        "login": {"login", "Логин", "char(64)", false},
    }
    sessionModel.Fields = tmp
    return sessionModel
}

type SessionModel struct {
    Entity
}

type Session struct {
    Sid       string     `json:"sid"`
    Login     string     `json:"login"`
    TableData *UserModel `json:"tableData"`
}
