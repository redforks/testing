package testdb_test

import (
	"github.com/redforks/testing/reset"

	"gopkg.in/mgo.v2"

	. "github.com/redforks/testing/mongo/testdb"

	bdd "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Rec struct {
	Id   int `bson:"_id"`
	Name string
}

const testDBName = "unittest"
const testDBUrl = "/" + testDBName

var _ = bdd.Describe("mongotest", func() {
	var (
		testDb *TestDb
	)

	bdd.BeforeEach(func() {
		reset.Enable()
		testDb = New(testDBUrl)
	})

	bdd.AfterEach(func() {
		reset.Disable()

		session, err := mgo.Dial("127.0.0.1")
		Ω(err).Should(Succeed())
		defer session.Close()

		// ensure database are deleted in reset
		dbs, err := session.DatabaseNames()
		Ω(err).Should(Succeed())
		Ω(dbs).ShouldNot(ContainElement(testDBName))
	})

	bdd.It("Fields", func() {
		Ω(testDb.Session).ShouldNot(BeNil())
		Ω(testDb.Database).ShouldNot(BeNil())
		Ω(testDb.Name).Should(Equal(testDBName))
		Ω(testDb.Url()).Should(Equal(testDBUrl))
	})

	bdd.It("Insert", func() {
		recs := []*Rec{
			{0, "foo"}, {1, "bar"}, {2, "foobar"},
		}
		Ω(testDb.C("foo").Insert(recs[0], recs[1], recs[2])).Should(Succeed())
		back := []*Rec{}
		Ω(testDb.C("foo").Find(nil).All(&back)).Should(Succeed())
		Ω(back).Should(Equal(recs))
	})

})
