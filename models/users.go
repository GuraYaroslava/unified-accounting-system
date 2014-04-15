package models

func (c *ModelManager) Users() *UserModel {
    userModel := new(UserModel)

    userModel.TableName = "Users"
    userModel.Caption = "Пользователи"

    //admin mode
    userModel.Columns = []string{"id", "login", "password", "salt", "sid", "fname", "lname", "pname", "email", "phone", "address"}
    userModel.ColNames = []string{"Id", "Логин", "Хеш", "Соль", "Sid", "Фамилия", "Имя", "Отчество", "E-mail", "Телефон", "Адрес"}

    //user mode
    userModel.UserColumns = []string{"id", "fname", "lname", "pname", "email", "phone", "address"}
    userModel.UserColNames = []string{"Id", "Фамилия", "Имя", "Отчество", "E-mail", "Телефон", "Адрес"}

    tmp := map[string]*Field{
        "id":       {"id", "Id", "serial", false},
        "login":    {"login", "Логин", "varchar(32)", false},
        "password": {"password", "Хеш", "varchar(128)", false},
        "salt":     {"salt", "Соль", "varchar(64)", false},
        "sid":      {"sid", "Sid", "varchar(40)", false},

        "fname": {"fname", "Фамилия", "varchar(32)", false},
        "lname": {"lname", "Имя", "varchar(32)", false},
        "pname": {"pname", "Отчество", "varchar(32)", false},

        "email":   {"email", "E-mail", "varchar(32)", false},
        "phone":   {"phone", "Телефон", "varchar(32)", false},
        "address": {"address", "Адрес", "varchar(32)", false},
    }
    userModel.Fields = tmp
    return userModel
}

type UserModel struct {
    Entity
    UserColumns  []string
    UserColNames []string
}
