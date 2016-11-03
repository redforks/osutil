package osutil_test

import (
	. "github.com/redforks/osutil"
	"io/ioutil"
	"path/filepath"

	"github.com/redforks/testing/iotest"
	"github.com/redforks/testing/reset"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("WriteFile", func() {

	var (
		testDir iotest.TempTestDir
	)

	BeforeEach(func() {
		reset.Enable()
		testDir = iotest.NewTempTestDir()
	})

	AfterEach(func() {
		reset.Disable()
	})

	It("File not exist", func() {
		f := filepath.Join(testDir.Dir(), "foo")
		Ω(WriteFile(f, []byte("foobar"), 0700, 0777)).Should(Succeed())
		Ω(ioutil.ReadFile(f)).Should(BeEquivalentTo("foobar"))
	})

	It("file exist", func() {
		f := filepath.Join(testDir.Dir(), "foo")
		Ω(ioutil.WriteFile(f, []byte("bar"), 0777)).Should(Succeed())
		Ω(WriteFile(f, []byte("foobar"), 0700, 0777)).Should(Succeed())
		Ω(ioutil.ReadFile(f)).Should(BeEquivalentTo("foobar"))
	})

	It("Directory not exist", func() {
		f := filepath.Join(testDir.Dir(), "abc/cde/foo")
		Ω(WriteFile(f, []byte("foobar"), 0700, 0777)).Should(Succeed())
		Ω(ioutil.ReadFile(f)).Should(BeEquivalentTo("foobar"))
	})

})
