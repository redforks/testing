package testdb_test

import (
	"testing"

	. "github.com/redforks/testing/mongo/testdb"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

type Rec struct {
	Id   int `bson:"_id"`
	Name string
}

const testDBName = "unittest"
const testDBUrl = "/" + testDBName

func ensureDBDeleted(t *testing.T) {
	session, err := mgo.Dial("127.0.0.1")
	assert.NoError(t, err)
	defer session.Close()

	// ensure database are deleted in reset
	dbs, err := session.DatabaseNames()
	assert.NoError(t, err)
	assert.NotContains(t, testDBName, dbs)
}

func TestFields(t *testing.T) {
	testDb := New(testDBUrl)
	defer testDb.Close()

	assert.NotNil(t, testDb.Session)
	assert.NotNil(t, testDb.Database)
	assert.Equal(t, testDBName, testDb.Name)
	assert.Equal(t, testDBUrl, testDb.Url())

	ensureDBDeleted(t)
}

func TestInsert(t *testing.T) {
	testDb := New(testDBUrl)
	defer testDb.Close()

	recs := []*Rec{
		{0, "foo"}, {1, "bar"}, {2, "foobar"},
	}
	assert.NoError(t, testDb.C("foo").Insert(recs[0], recs[1], recs[2]))
	back := []*Rec{}
	assert.NoError(t, testDb.C("foo").Find(nil).All(&back))
	assert.Equal(t, recs, back)

	ensureDBDeleted(t)
}
