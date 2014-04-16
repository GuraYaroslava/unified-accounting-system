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