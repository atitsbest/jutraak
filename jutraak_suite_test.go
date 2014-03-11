package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestJutraak(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jutraak Suite")
}
