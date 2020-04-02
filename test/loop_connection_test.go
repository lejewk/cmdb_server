package test

import (
    "github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
    "testing"
    "time"
)

func TestLoopConnection(t *testing.T) {
    for i := 0; i < 100; i++ {
        mgo.Dial("localhost")
    }
}

func TestSession(t *testing.T) {
    s, _ := mgo.Dial("localhost")

    go func() {
        s.DB("cmdb").C("data_source").Find(bson.M{"_id": "1"})
        time.Sleep(5 * time.Second)
    }()

    go func() {
        s.DB("cmdb").C("data_source").Find(bson.M{"_id": "2"})
        time.Sleep(5 * time.Second)
    }()
}

func TestLoopConnectionPool(t *testing.T) {
    for i := 0; i < 100; i++ {
        mgo.DialWithInfo(&mgo.DialInfo{
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
    }
}