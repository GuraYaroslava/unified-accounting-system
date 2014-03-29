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

func DBSelect(from, where string, fields ...string) string {
    var format string = "SELECT %s FROM %s"
    if where != "" {
        return fmt.Sprintf(format+" WHERE %s", strings.Join(fields, ", "), from, where)
    } else {
        return fmt.Sprintf(format, strings.Join(fields, ", "), from)
    }
}

func DBInsert(into string, fields []string) string {
    var format string = "INSERT INTO %s (%s) VALUES (%s);"
    return fmt.Sprintf(format, into, strings.Join(fields, ", "), strings.Join(MakeParams(fields), ", "))
}

func DBUpdate(table string, fields []string, where string) string {
    var format string = "UPDATE %s SET %s WHERE %s"
    return fmt.Sprintf(format, table, strings.Join(MakePair(fields), ", "), where)
}

func MakeParams(fields []string) []string {
    var result = make([]string, len(fields))
    for i := 0; i < len(fields); i++ {
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
