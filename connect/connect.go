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
    query = `CREATE OR REPLACE FUNCTION del_elem_by_index(anyarray, integer)
              RETURNS anyarray AS
            $BODY$ 
                declare 
                    arr_ $1%type;
                    idx_ $2%type;
                begin
                    for idx_ in array_lower($1, 1)..array_upper($1, 1) loop
                        if not idx_ = $2 then 
                            arr_ = array_append(arr_, $1[idx_]);
                        else 
                            arr_ = array_append(arr_, NULL);
                        end if;
                    end loop;
                    return arr_;
                end;
            $BODY$
              LANGUAGE plpgsql VOLATILE
            COST 100;
            ALTER FUNCTION del_elem_by_index(anyarray, integer)
            OWNER TO admin;`
    _, err = db.Exec(query)
    utils.HandleErr("[Connect.InitSchema.del_elem_by_index]: ", err)

    query = `CREATE OR REPLACE FUNCTION del_null_from_arr(anyarray)
              RETURNS anyarray AS
            $BODY$ 
                declare
                    arr_ $1%type;
                    idx_ int;
                begin
                    for idx_ in array_lower($1, 1)..array_upper($1, 1) loop
                        if not $1[idx_] = '' then
                            arr_ = array_append(arr_, $1[idx_]);
                        end if;
                    end loop;
                    return arr_;
                end;
            $BODY$
              LANGUAGE plpgsql VOLATILE
                COST 100;
            ALTER FUNCTION del_null_from_arr(anyarray)
                  OWNER TO admin;`
    _, err = db.Exec(query)
    utils.HandleErr("[Connect.InitSchema.del_null_from_arr]: ", err)
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

func DBGetColumnNames(tableName string) []string {
    db := DBConnect()
    defer DBClose(db)
    var columns []string
    query := "SELECT * FROM " + tableName
    stmt, err := db.Prepare(query)
    utils.HandleErr("[Connect.DBGetColumnNames] Prepare: ", err)
    rows, err := stmt.Query()
    utils.HandleErr("[Connect.DBGetColumnNames] Query: ", err)
    columns, err = rows.Columns()
    utils.HandleErr("[Connect.DBGetColumnNames] Columns: ", err)
    fmt.Println("column names: ", columns)
    return columns
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
