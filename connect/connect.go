package connect

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "github.com/uas/utils"
    "strconv"
    "strings"
)

const user string = "admin"
const dbname string = "uas"
const password string = "admin"

type DBComps interface {
    Close() error
}

func DBConnect() *sql.DB {
    db, err := sql.Open("postgres", "host=localhost user="+user+" dbname="+dbname+" password="+password+" sslmode=disable")
    utils.HandleErr("Coonect DB: ", err)
    return db
}

func DBClose(comps ...DBComps) {
    for _, comp := range comps {
        if comp != nil {
            comp.Close()
        }
    }
}

func DBInitSchema() {
    db := DBConnect()
    defer DBClose(db)
    query := `CREATE TABLE IF NOT EXISTS users (
                    id       serial       NOT NULL PRIMARY KEY,
                    login    varchar(32)  NOT NULL UNIQUE,
                    password varchar(128) NOT NULL,
                    salt     varchar(64)  NOT NULL,
                    sid      varchar(40)  NOT NULL DEFAULT '',
                
                    fname    varchar(32)  NOT NULL DEFAULT '',
                    lname    varchar(32)  NOT NULL DEFAULT '',
                    pname    varchar(32)  NOT NULL DEFAULT '',
                    email    varchar(32)  NOT NULL DEFAULT '',
                    phone    varchar(32)  NOT NULL DEFAULT '',
                
                    region   varchar(32)  NOT NULL DEFAULT '',
                    district varchar(32)  NOT NULL DEFAULT '',
                    city     varchar(32)  NOT NULL DEFAULT '',
                    street   varchar(32)  NOT NULL DEFAULT '',
                    building varchar(32)  NOT NULL DEFAULT ''
                );`
    _, err := db.Exec(query)
    utils.HandleErr("[Connect.InitSchema.Users]: ", err)
    query = `CREATE TABLE IF NOT EXISTS contests (
                    id       serial       NOT NULL PRIMARY KEY,
                    name     varchar(128) NOT NULL UNIQUE,
                    date     date         NOT NULL
                );`
    _, err = db.Exec(query)
    utils.HandleErr("[Connect.InitSchema.Contests]: ", err)
    query = `CREATE TABLE IF NOT EXISTS blanks (
                    id          serial        NOT NULL PRIMARY KEY,
                    name        varchar(128)  NOT NULL UNIQUE,
                    contest_id  int           NOT NULL REFERENCES contests (id) ON DELETE CASCADE,
                    columns     varchar(32)[],
                    colNames    varchar(32)[],
                    types       varchar(32)[]
                );`
    _, err = db.Exec(query)
    utils.HandleErr("[Connect.InitSchema.Contests]: ", err)
}

func DBGetLastInsertedId(tableName string) string {
    db := DBConnect()
    defer DBClose(db)
    var id string
    query := "SELECT last_value FROM " + tableName + "_id_seq"
    stmt, err := db.Prepare(query)
    utils.HandleErr("[Connect.GetLastInsertedId] Prepare: ", err)
    err = stmt.QueryRow().Scan(&id)
    utils.HandleErr("[Connect.GetLastInsertedId] Query: ", err)
    return id
}

func DBSelect(from string, where []string, fields ...string) string {
    var format string = "SELECT %s FROM %s"
    if len(where) > 0 {
        return fmt.Sprintf(format+" WHERE %s", strings.Join(fields, ", "), from, strings.Join(MakePair(where), ", "))
    } else {
        return fmt.Sprintf(format, strings.Join(fields, ", "), from)
    }
}

func DBInsert(into string, fields []string) string {
    var format string = "INSERT INTO %s (%s) VALUES (%s);"
    return fmt.Sprintf(format, into, strings.Join(fields, ", "), strings.Join(MakeParams(len(fields)), ", "))
}

func DBUpdate(table string, fields []string, where string) string {
    var format string = "UPDATE %s SET %s WHERE %s"
    return fmt.Sprintf(format, table, strings.Join(MakePair(fields), ", "), where)
}

func DBDelete(table string, field string, params []interface{}) string {
    var format string = "DELETE FROM %s WHERE %s IN (%s)"
    return fmt.Sprintf(format, table, field, strings.Join(MakeParams(len(params)), ", "))
}

func MakeParams(n int) []string {
    var result = make([]string, n)
    for i := 0; i < n; i++ {
        result[i] = "$" + strconv.Itoa(i+1)
    }
    return result
}

func MakePair(fields []string) []string {
    var result = make([]string, len(fields))
    for i := 0; i < len(fields); i++ {
        result[i] = fields[i] + "=$" + strconv.Itoa(i+1)
    }
    return result
}
