package goactivedirectory_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	activedirectory "github.com/gaikwadpratik/go-activedirectory"
)

var _ = Describe("RootDse", func() {
	It("Should get Entries for Root configuration", func() {
		serverConfig := &activedirectory.ServerConfig{
			Url: os.Getenv("ldapUrl"),
		}
		entries, err := activedirectory.GetRootDSE(serverConfig)
		Expect(err).Should(BeNil(), "Error is not nil while getting root dse entries")
		Expect(entries).ShouldNot(BeEmpty(), "No entries received")
	})

	It("Should get DefaultNamingContext of Root server", func() {
		serverConfig := &activedirectory.ServerConfig{
			Url: os.Getenv("ldapUrl"),
		}
		val, err := activedirectory.GetAttributeOnRootDSE(serverConfig, "defaultNamingContext")
		Expect(err).Should(BeNil(), "Error is not nil while getting default naming context")
		Expect(val).ShouldNot(BeEmpty(), "Default naming context is not returned")
		Expect(val).Should(Equal(os.Getenv("DefaultNamingContext")), "Different default naming context received: %s", val)
	})
})
