package testdb

import (
	"github.com/redforks/testing/reset"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

// TestDb create a connection to mongodb, and delete the database in reset. Also
// provide mongodb operation methods, so we do not need to deal with returned error.
// *mgo.Databae is nested in, so all Database method con be used directly from TestDb, such as:
//
//   db := testdb.New("blah_test")
//   db.C("tbl").Insert(...
//
// Instead of:
//
//   db.Session.DB("").C("tbl").Insert(...
type TestDb struct {
	Session *mgo.Session
	*mgo.Database
	url string
}

// Create a new TestDb instance, name of the test database, if empty, uses
// "test". Recommand use a named database, because after each test the database
// will be deleted, pass a unique name to prevent break others. TestDb always
// connect to 127.0.0.1.
func New(t assert.TestingT, name string) *TestDb {
	url := "127.0.0.1/" + name
	session, err := mgo.Dial(url)
	assert.NoError(t, err)
	reset.Add(func() {
		assert.NoError(t, session.DB("").DropDatabase())
		session.Close()
	})
	return &TestDb{session, session.DB(""), url}
}

// Return mongodb connection url
func (db *TestDb) Url() string {
	return db.url
}
