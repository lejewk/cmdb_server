package factory

import (
    "github.com/globalsign/mgo"
    "log"
    "sync"
)

var once sync.Once
var session *mgo.Session

type MongoDb struct { }

func (MongoDb) Session() *mgo.Session {
    once.Do(func() {
        s, err := mgo.DialWithInfo(&mgo.DialInfo{
            Addrs:          []string{"localhost"},
            Timeout:        0,
            Database:       "cmdb",
            PoolLimit:      20,
            PoolTimeout:    0,
            ReadTimeout:    0,
            WriteTimeout:   0,
            MinPoolSize:    10,
            MaxIdleTimeMS:  15,
        })

        if err != nil {
            log.Panic("Failed connect to mongo db")
        }

        session = s
    })

    return session
}
