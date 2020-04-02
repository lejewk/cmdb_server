package test

import (
    "cmdb_server/model"
    "cmdb_server/repository"
    "log"
    "testing"
)

var (
    testMallId = "ectlocal"
    testKey = "postgresql_mall_master"
)

func TestInsertOrUpdate(t *testing.T) {
    result := repository.GetDataSourceRepository().InsertOrUpdate(
        testMallId,
        testKey,
        model.DatabaseConnection{
            Host: "localhost",
            Port: 9999,
            Database: "mall",
            Username: "malluser",
            Password: "password",
        })

    result = repository.GetDataSourceRepository().InsertOrUpdate(
        testMallId,
        "postgresl_mall_slave",
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
    result := repository.GetDataSourceRepository().FindById(testMallId)
    log.Println(result)
}

func TestDeleteByIdAndKey(t *testing.T) {
    result := repository.GetDataSourceRepository().DeleteByIdAndKey(testMallId, testKey)
    log.Println(result)
}
