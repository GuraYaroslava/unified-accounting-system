package models

func (c *ModelManager) Users() *UserModel {
	userModel := new(UserModel)

	userModel.TableName = "Users"
	userModel.Caption = "Пользователи"
	userModel.Columns = []string{"id", "password", "fname", "lname", "pname", "login", "salt", "hash", "email", "phone", "address"}
	tmp := map[string]*Field{
		"id":      {"id", "Id", "integer", false},
		"fname":   {"fname", "Фамилия", "char(64)", false},
		"lname":   {"lname", "Имя", "char(64)", false},
		"pname":   {"pname", "Отчество", "char(64)", false},
		"login":   {"login", "Логин", "char(64)", false},
		"salt":    {"salt", "Соль", "char(64)", false},
		"hash":    {"hash", "Хеш", "char(64)", false},
		"email":   {"email", "E-mail", "char(64)", false},
		"phone":   {"phone", "Телефон", "char(64)", false},
		"address": {"address", "Адрес", "char(64)", false},
	}
	userModel.Fields = tmp
	return userModel
}

type UserModel struct {
	Entity
}

type User struct {
	Id        string     `json:"id"`
	FName     string     `json:"fname"`
	LName     string     `json:"lname"`
	PName     string     `json:"pname"`
	Login     string     `json:"login"`
	Salt      string     `json:"salt"`
	Hash      string     `json:"hash"`
	EMail     string     `json:"email"`
	Phone     string     `json:"phone"`
	Address   string     `json:"address"`
	TableData *UserModel `json:"tableData"`
}
