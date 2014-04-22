package models

func (c *ModelManager) Users() *UserModel {
    userModel := new(UserModel)

    userModel.TableName = "Users"
    userModel.Caption = "Пользователи"

    //admin mode
    userModel.Columns = []string{
        "id",
        "login",
        "password",
        "salt",
        "sid",
        "fname",
        "lname",
        "pname",
        "email",
        "phone",
        "region",
        "district",
        "city",
        "street",
        "building"}
    userModel.ColNames = []string{
        "Id",
        "Логин",
        "Хеш",
        "Соль",
        "Sid",
        "Фамилия",
        "Имя",
        "Отчество",
        "E-mail",
        "Телефон",
        "Регион",
        "Район",
        "Город",
        "Улица",
        "Дом/кв."}

    //user mode-----------------------------------------------------------------
    userModel.UserColumns = []string{
        "id",
        "fname",
        "lname",
        "pname",
        "email",
        "phone",
        "region",
        "district",
        "city",
        "street",
        "building"}
    userModel.UserColNames = []string{
        "Id",
        "Фамилия",
        "Имя",
        "Отчество",
        "E-mail",
        "Телефон",
        "Регион",
        "Район",
        "Город",
        "Улица",
        "Дом/кв."}

    userModel.UserTypes = []string{
        "serial",
        "varchar(32)",
        "varchar(32)",
        "varchar(32)",
        "varchar(32)",
        "varchar(32)",
        "varchar(32)",
        "varchar(32)",
        "varchar(32)",
        "varchar(32)",
        "varchar(32)"}
    //--------------------------------------------------------------------------

    tmp := map[string]*Field{
        "id":       {"id", "Id", "serial NOT NULL PRIMARY KEY", false},
        "login":    {"login", "Логин", "varchar(32) NOT NULL UNIQUE", false},
        "password": {"password", "Хеш", "varchar(128)", false},
        "salt":     {"salt", "Соль", "varchar(64)", false},
        "sid":      {"sid", "Sid", "varchar(40)", false},

        "fname": {"fname", "Фамилия", "varchar(32) NOT NULL DEFAULT ''", false},
        "lname": {"lname", "Имя", "varchar(32) NOT NULL DEFAULT ''", false},
        "pname": {"pname", "Отчество", "varchar(32) NOT NULL DEFAULT ''", false},

        "email": {"email", "E-mail", "varchar(32) NOT NULL DEFAULT ''", false},
        "phone": {"phone", "Телефон", "varchar(32) NOT NULL DEFAULT ''", false},

        "region":   {"region", "Регион", "varchar(32) NOT NULL DEFAULT ''", false},
        "district": {"district", "Район", "varchar(32) NOT NULL DEFAULT ''", false},
        "city":     {"city", "Город", "varchar(32) NOT NULL DEFAULT ''", false},
        "street":   {"street", "Улица", "varchar(32) NOT NULL DEFAULT ''", false},
        "building": {"building", "Дом/кв.", "varchar(32) NOT NULL DEFAULT ''", false},
    }
    userModel.Fields = tmp
    return userModel
}

type UserModel struct {
    Entity
    UserColumns  []string
    UserColNames []string
    UserTypes    []string
}
