package connect

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "github.com/uas/utils"
    "strings"
)

const user string = "admin"
const dbname string = "acs"
const password string = "admin"

type DBComps interface {
    Close() error
}

func DBConnect() *sql.DB {
    db, err := sql.Open("postgres", "user="+user+" dbname="+dbname+" password="+password+" sslmode=disable")
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

func DBSelect(from, where string, fields ...string) string {
    var format string = "SELECT %s FROM %s"
    if where != "" {
        return fmt.Sprintf(format+" WHERE %s", strings.Join(fields, ", "), from, where)
    } else {
        return fmt.Sprintf(format, strings.Join(fields, ", "), from)
    }
}

func DBInsert(into string, fields, params []string) string {
    var format string = "INSERT INTO %s (%s) VALUES (%s);"
    return fmt.Sprintf(format, into, strings.Join(fields, ", "), strings.Join(params, ", "))
}
