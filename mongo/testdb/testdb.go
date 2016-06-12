package testdb

import (
	"github.com/redforks/testing/reset"

	. "github.com/onsi/gomega"
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

// Create a new TestDb instance.
func New(dbUrl string) *TestDb {
	session, err := mgo.Dial(dbUrl)
	Ω(err).Should(Succeed())
	reset.Add(func() {
		Ω(session.DB("").DropDatabase()).Should(Succeed())
		session.Close()
	})
	return &TestDb{session, session.DB(""), dbUrl}
}

// Return mongodb connection url
func (db *TestDb) Url() string {
	return db.url
}
