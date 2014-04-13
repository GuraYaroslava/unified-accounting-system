package models

import (
    "fmt"
    "github.com/uas/connect"
    "github.com/uas/utils"
    "time"
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
    ColNames  []string
}

func (this Entity) Select(where map[string]interface{}, fields ...string) []interface{} {
    keys := make([]string, len(where))
    vals := make([]interface{}, len(where))
    j := 0
    for i, v := range where {
        keys[j] = i
        vals[j] = v
        j++
    }
    db := connect.DBConnect()
    query := connect.DBSelect(this.TableName, keys, fields...)
    stmt, err := db.Prepare(query)
    utils.HandleErr("[Entity.Select] Prepare: ", err)
    defer connect.DBClose(db, stmt)

    rows, err := stmt.Query(vals...)
    utils.HandleErr("[Entity.Select] Query: ", err)

    rowsInf, err := stmt.Exec(vals...)
    utils.HandleErr("[Entity.Select] Exec: ", err)

    columns, err := rows.Columns()
    row := make([]interface{}, len(columns))
    values := make([]interface{}, len(columns))
    for i, _ := range row {
        row[i] = &values[i]
    }

    l, err := rowsInf.RowsAffected()
    utils.HandleErr("[Entity.Select] RowsAffected: ", err)
    answer := make([]interface{}, l)
    j = 0

    for rows.Next() {
        rows.Scan(row...)
        answer[j] = make(map[string]interface{}, len(values))
        record := make(map[string]interface{}, len(values))
        for i, col := range values {
            if col != nil {
                //fmt.Printf("\n%s: type= %s\n", columns[i], reflect.TypeOf(col))
                switch col.(type) {
                default:
                    utils.HandleErr("Entity.Select: Unexpected type.", nil)
                case bool:
                    record[columns[i]] = col.(bool)
                case int:
                    record[columns[i]] = col.(int)
                case int64:
                    record[columns[i]] = col.(int64)
                case float64:
                    record[columns[i]] = col.(float64)
                case string:
                    record[columns[i]] = col.(string)
                case []byte:
                    record[columns[i]] = string(col.([]byte))
                case time.Time:
                    record[columns[i]] = col
                }
            }
            answer[j] = record
        }
        j++
    }
    return answer
}

func (this Entity) Insert(fields []string, params []interface{}) {
    db := connect.DBConnect()
    query := connect.DBInsert(this.TableName, fields)
    stmt, err := db.Prepare(query)
    utils.HandleErr("[Entity.Insert] Prepare: ", err)
    defer connect.DBClose(db, stmt)
    _, err = stmt.Exec(params...)
    utils.HandleErr("[Entity.Insert] Exec: ", err)
}

func (this Entity) Update(fields []string, params []interface{}, where string) {
    db := connect.DBConnect()
    query := connect.DBUpdate(this.TableName, fields, where)
    fmt.Println("Update: ", query)
    stmt, err := db.Prepare(query)
    utils.HandleErr("[Entity.Update] Prepare: ", err)
    defer connect.DBClose(db, stmt)
    _, err = stmt.Exec(params...)
    utils.HandleErr("[Entity.Update] Exec: ", err)
}

func (this Entity) Delete(field string, params []interface{}) {
    db := connect.DBConnect()
    query := connect.DBDelete(this.TableName, field, params)
    stmt, err := db.Prepare(query)
    utils.HandleErr("[Entity.Delete] Prepare: ", err)
    defer connect.DBClose(db, stmt)
    _, err = stmt.Exec(params...)
    utils.HandleErr("[Entity.Delete] Exec: ", err)
}
