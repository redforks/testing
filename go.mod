module github.com/redforks/testing

go 1.12

require (
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	github.com/stretchr/testify v1.3.0
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
)

replace gopkg.in/mgo.v2 => github.com/redforks/mgo v0.0.0-20170322165704-f51d5a76a374
