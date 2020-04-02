package test

import (
    "github.com/globalsign/mgo"
    "testing"
)

func TestLoopConnection(t *testing.T) {
    for i := 0; i < 100; i++ {
        mgo.Dial("localhost")
    }
}
