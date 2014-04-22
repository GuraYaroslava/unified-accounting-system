DROP DATABASE IF EXISTS "uas";

CREATE DATABASE "uas";

CREATE USER admin WITH password "admin";

GRANT ALL ON DATABASE "uas" TO admin;

CREATE TABLE IF NOT EXISTS users (
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

    region   varchar(32)  NOT NULL DEFAULT ''
    district varchar(32)  NOT NULL DEFAULT ''
    city     varchar(32)  NOT NULL DEFAULT ''
    street   varchar(32)  NOT NULL DEFAULT ''
    building varchar(32)  NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS contests (
    id       serial       NOT NULL PRIMARY KEY,
    name     varchar(128) NOT NULL UNIQUE,
    date     date         NOT NULL
);

CREATE TABLE IF NOT EXISTS blanks (
    id         serial       NOT NULL PRIMARY KEY,
    name       varchar(128) NOT NULL UNIQUE,
    contest_id int          NOT NULL REFERENCES contests (id) ON DELETE CASCADE,
    columns    varchar(32)[],
    colNames   varchar(32)[],
    types      varchar(32)[]
);

CREATE OR REPLACE FUNCTION del_elem_by_index(anyarray, integer)
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
        OWNER TO admin;

CREATE OR REPLACE FUNCTION del_null_from_arr(anyarray)
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
        OWNER TO admin;

