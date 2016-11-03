package osutil_test

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"testing"
)

func TestOsutil(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Osutil Suite")
}
