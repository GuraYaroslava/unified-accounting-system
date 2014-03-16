package models

import (
    "database/sql"
    "github.com/hope/connect"
    "github.com/hope/utils"
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
}

func (this Entity) Select(where string, fields ...string) sql.Result {
    db := connect.DBConnect()
    stmt := connect.DBSelect(this.TableName, where, fields...)
    rows, err := db.Exec(stmt)
    utils.HandleErr("[Select] Exec: ", err)
    return rows
}
