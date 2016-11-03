package osutil_test

import (
	"crypto/rand"
	. "github.com/redforks/osutil"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/redforks/testing/iotest"
	"github.com/redforks/testing/reset"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("File", func() {
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

	Context("Copy", func() {
		var (
			contents []byte

			srcFile, dstFile string
		)

		BeforeEach(func() {
			srcFile = filepath.Join(testDir.Dir(), "a")
			contents = make([]byte, 5000)
			Ω(rand.Read(contents)).Should(Equal(5000))
			Ω(ioutil.WriteFile(srcFile, contents, 0600)).Should(Succeed())
		})

		AfterEach(func() {
			Ω(ioutil.ReadFile(dstFile)).Should(Equal(contents))
		})

		It("Dest file not exist", func() {
			dstFile = filepath.Join(testDir.Dir(), "b")
			Ω(Copy(dstFile, srcFile)).Should(Succeed())
		})

		It("Dest is a directory", func() {
			dst := filepath.Join(testDir.Dir(), "b")
			dstFile = filepath.Join(dst, "a")
			Ω(os.Mkdir(dst, 0700)).Should(Succeed())
			Ω(Copy(dst, srcFile)).Should(Succeed())
		})

		It("Overwrite dest file", func() {
			dstFile = filepath.Join(testDir.Dir(), "b")
			Ω(ioutil.WriteFile(dstFile, []byte("foo"), 0600)).Should(Succeed())
			Ω(Copy(dstFile, srcFile)).Should(Succeed())
		})

	})
})
