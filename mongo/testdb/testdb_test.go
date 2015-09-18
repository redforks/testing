package testdb

import (
	"gopkg.in/mgo.v2"

	bdd "github.com/onsi/ginkgo"
	"github.com/redforks/testing/reset"
	"github.com/stretchr/testify/assert"
)

type Rec struct {
	Id   int "_id"
	Name string
}

var _ = bdd.Describe("mongotest", func() {
	const (
		DBNAME = "mongotest_test"
	)

	var (
		testDb *TestDb
	)

	bdd.BeforeEach(func() {
		reset.Enable()
		testDb = New(t(), DBNAME)
	})

	bdd.AfterEach(func() {
		reset.Disable()

		session, err := mgo.Dial("127.0.0.1")
		assert.NoError(t(), err)
		defer session.Close()

		// ensure database are deleted in reset
		dbs, err := session.DatabaseNames()
		assert.NoError(t(), err)
		assert.NotContains(t(), dbs, DBNAME)
	})

	bdd.It("Fields", func() {
		assert.NotNil(t(), testDb.Session)
		assert.NotNil(t(), testDb.Database)
		assert.Equal(t(), DBNAME, testDb.Name)
		assert.Equal(t(), "127.0.0.1/"+DBNAME, testDb.Url())
	})

	bdd.It("Insert", func() {
		recs := []*Rec{
			{0, "foo"}, {1, "bar"}, {2, "foobar"},
		}
		assert.NoError(t(), testDb.C("foo").Insert(recs[0], recs[1], recs[2]))
		back := []*Rec{}
		assert.NoError(t(), testDb.C("foo").Find(nil).All(&back))
		assert.Equal(t(), recs, back)
	})

})
