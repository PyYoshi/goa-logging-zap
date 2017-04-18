package goazap_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestZap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Zap Suite")
}
