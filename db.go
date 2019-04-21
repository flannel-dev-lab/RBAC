package RBAC

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var (
    DbPoolReady bool

    DBRead  *sql.DB
    DBWrite *sql.DB
)

func DbInit() {
    if !DbPoolReady {
        LoadConfiguration()
        ConnectMysqlPool()
    
        DbPoolReady = true
    }
}

func ConnectMysqlPool() {
    reader, err := sql.Open("mysql", Config.DBReader.Username+":"+Config.DBReader.Password+"@tcp("+Config.DBReader.Host+":"+Config.DBReader.Port+")/"+Config.DBReader.Name)

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    DBRead = reader

    writer, err := sql.Open("mysql", Config.DBWriter.Username+":"+Config.DBWriter.Password+"@tcp("+Config.DBWriter.Host+":"+Config.DBWriter.Port+")/"+Config.DBWriter.Name)

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    DBWrite = writer

    return;
}
