package repository

import (
    "cmdb_server/factory"
    "cmdb_server/model"
    "github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
    "log"
    "sync"
)

var (
    database   = "cmdb"
    collection = "data_source"
)

type dataSourceRepository struct {
    session *mgo.Session
}

var once sync.Once
var instance *dataSourceRepository

func GetDataSourceRepository() *dataSourceRepository {
    once.Do(func() {
        instance = &dataSourceRepository{session: factory.MongoDb{}.Session()}
    })
    return instance
}

func (s *dataSourceRepository) coll() *mgo.Collection {
    return s.session.DB(database).C(collection)
}

func (s *dataSourceRepository) InsertOrUpdate(mallId string, key string, update model.DatabaseConnection) *mgo.ChangeInfo {
    result, err := s.coll().Upsert(
        bson.M{"_id": mallId},
        bson.M{"_id": mallId, key: update},
    )

    if err != nil {
        log.Fatal("Failed insert or update", err)
    }

    return result
}

func (s *dataSourceRepository) FindById(mallId string) *bson.M {
    var data bson.M
    err := s.coll().FindId(mallId).One(&data)

    if err != nil {
        log.Fatal("Failed find id & all", err)
    }

    return &data
}
