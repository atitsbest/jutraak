package repositories

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SqlProblemRepository", func() {
  var (
    sut *SqlProblemRepository
    err error
  )

  Describe("Connect to database", func() {
    Context("Valid connection string", func() {
      sut, err = NewSqlProblemRepsoitory("right")

      It("should create repository", func() {
        Expect(err).To(Equal(nil))
        Expect(sut).NotTo(Equal(nil))
      })
    })

    Context("Invalid connection string", func() {
      sut, err = NewSqlProblemRepsoitory("wrong")

      It("should return an error an no repository", func() {
        Expect(err).NotTo(Equal(nil))
        Expect(sut).To(Equal(nil))
      })
    })
  })
})

