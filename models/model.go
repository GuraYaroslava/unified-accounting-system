package models

import (
    "database/sql"
    "fmt"
    "github.com/uas/connect"
    "github.com/uas/utils"
)

type ModelManager struct{}

type Field struct {
    Name    string
    Caption string
    Type    string
    Ref     bool
}

type Entity struct {
    TableName string
    Caption   string
    Fields    map[string]*Field
    Columns   []string
}

func (this Entity) Select(where string, fields ...string) sql.Result {
    db := connect.DBConnect()
    query := connect.DBSelect(this.TableName, where, fields...)
    result, err := db.Exec(query)
    utils.HandleErr("[Select] Exec: ", err)
    return result
}

func (this Entity) Insert(into string, fields []string, params []string) sql.Result {
    db := connect.DBConnect()
    query := connect.DBInsert(this.TableName, fields, params)
    result, err := db.Exec(query)
    utils.HandleErr("[Select] Exec: ", err)
    return result
}
