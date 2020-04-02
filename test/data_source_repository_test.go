package test

import (
    "cmdb_server/model"
    "cmdb_server/repository"
    "log"
    "testing"
)

func TestInsertOrUpdate(t *testing.T) {
    result := repository.GetDataSourceRepository().InsertOrUpdate(
        "helloworld",
        "postgresql_mall_master",
        model.DatabaseConnection{
            Host: "localhost",
            Port: 9999,
            Database: "mall",
            Username: "malluser",
            Password: "password",
        })

    log.Println(result.UpsertedId)
    log.Println(result.Updated)
    log.Println(result.Matched)
}

func TestFindById(t *testing.T) {
    result := repository.GetDataSourceRepository().FindById("helloworld")
    log.Println(result)
}
