package util

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ArchiveHandler", func() {
	Describe("GetMtaIDFromArchive", func() {
		var mtaArchiveFilePath, _ = filepath.Abs("../test_resources/commands/mtaArchive.mtar")
		Context("with valid mta archive", func() {
			It("should extract and return the id from deployment descriptor", func() {
				Expect(GetMtaIDFromArchive(mtaArchiveFilePath)).To(Equal("test"))
			})
		})
		Context("with valid mta archive and no deployment descriptor provided", func() {
			It("should return error", func() {
				mtaArchiveFilePath, _ = filepath.Abs("../test_resources/util/mtaArchiveNoDescriptor.mtar")
				_, err := GetMtaIDFromArchive(mtaArchiveFilePath)
				Expect(err).To(MatchError("Could not get MTA id from archive"))
			})
		})

		Context("with invalid mta archive", func() {
			BeforeEach(func() {
				os.Create("test.mtar")
				mtaArchiveFilePath, _ = filepath.Abs("test.mtar")
			})
			It("should return error for not a valid zip archive", func() {
				_, err := GetMtaIDFromArchive(mtaArchiveFilePath)
				Expect(err).To(MatchError("zip: not a valid zip file"))
			})
			AfterEach(func() {
				os.Remove(mtaArchiveFilePath)
			})
		})

		Context("with oversize descriptor size", func() {
			BeforeEach(func(){
				setDefaultMaxDescriptorSize(64)
			})
			It("must return error for file exceeds max size limit", func(){
				var mtaArchiveFilePath, _ = filepath.Abs("../test_resources/commands/mtaArchive.mtar")
				_, err := GetMtaIDFromArchive(mtaArchiveFilePath)
				Expect(err).To(MatchError("The size 109 of file META-INF/mtad.yaml exceeds max size limit 64"))
			})
			AfterEach(func() {
				os.Remove(mtaArchiveFilePath)
			})
		})

	})
})
