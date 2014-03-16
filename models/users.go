package models

func (c *ModelManager) Users() *UserModel {
    userModel := new(UserModel)

    userModel.TableName = "Users"
    userModel.Caption = "Пользователи"

    userModel.Fields["id"].Caption = "Id"
    userModel.Fields["id"].Name = "Id"
    userModel.Fields["id"].Type = "int"
    userModel.Fields["id"].Ref = false

    userModel.Fields["fname"].Caption = "Фамилия"
    userModel.Fields["fname"].Name = "fname"
    userModel.Fields["fname"].Type = "varchar"
    userModel.Fields["fname"].Ref = false

    userModel.Fields["lname"].Caption = "Имя"
    userModel.Fields["lname"].Name = "lname"
    userModel.Fields["lname"].Type = "varchar"
    userModel.Fields["lname"].Ref = false

    userModel.Fields["patronymic"].Caption = "Отчество"
    userModel.Fields["patronymic"].Name = "patronymic"
    userModel.Fields["patronymic"].Type = "varchar"
    userModel.Fields["patronymic"].Ref = false

    userModel.Fields["login"].Caption = "Логин"
    userModel.Fields["login"].Name = "login"
    userModel.Fields["login"].Type = "varchar"
    userModel.Fields["login"].Ref = false

    userModel.Fields["salt"].Caption = "Соль"
    userModel.Fields["salt"].Name = "salt"
    userModel.Fields["salt"].Type = "int"
    userModel.Fields["salt"].Ref = false

    userModel.Fields["hash"].Caption = "Хеш"
    userModel.Fields["hash"].Name = "hash"
    userModel.Fields["hash"].Type = "varchar"
    userModel.Fields["hash"].Ref = false

    userModel.Fields["email"].Caption = "E-mail"
    userModel.Fields["email"].Name = "email"
    userModel.Fields["email"].Type = "varchar"
    userModel.Fields["email"].Ref = false

    userModel.Fields["phone"].Caption = "Телефон"
    userModel.Fields["phone"].Name = "phone"
    userModel.Fields["phone"].Type = "varchar"
    userModel.Fields["phone"].Ref = false

    userModel.Fields["address"].Caption = "Адрес"
    userModel.Fields["address"].Name = "address"
    userModel.Fields["address"].Type = "varchar"
    userModel.Fields["address"].Ref = false

    return userModel
}

type UserModel struct {
    Entity
}
