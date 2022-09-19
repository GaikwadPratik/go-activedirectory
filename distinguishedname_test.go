package goactivedirectory_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Distinguishedname", func() {
	Context("Group", func() {
		It("Should return DN name for the group if provided as input", func() {
			dnName, err := adInstance.GetGroupDistinguishedName(os.Getenv("GroupDNName"))
			Expect(err).Should(BeNil(), "err is not nil while getting DN for group")
			Expect(dnName).ShouldNot(BeEmpty(), "Distinguished name of the group is empty")
			Expect(dnName).Should(Equal(os.Getenv("GroupDNName")), "DN name is different than expected: %s", dnName)
		})

		It("Should get DN name for the group", func() {
			dnName, err := adInstance.GetGroupDistinguishedName(os.Getenv("GroupName"))
			Expect(err).Should(BeNil(), "err is not nil while getting DN for group")
			Expect(dnName).ShouldNot(BeEmpty(), "Distinguished name of the group is empty")
			Expect(dnName).Should(Equal(os.Getenv("GroupDNName")), "DN name is different than expected: %s", dnName)
		})

		It("Should return error for the group does not exist", func() {
			dnName, err := adInstance.GetGroupDistinguishedName(os.Getenv("GroupNonexistantName"))
			Expect(err).ShouldNot(BeNil(), "err is nil while getting DN for group")
			Expect(dnName).Should(BeEmpty(), "Distinguished name of the group is not empty: %s", dnName)
		})
	})

	Context("User", func() {
		It("Should return DN name for user if provided as input", func() {
			dnName, err := adInstance.GetUserDistinguishedName(os.Getenv("UsernameDN"))
			Expect(err).Should(BeNil(), "err is not nil while getting DN for user")
			Expect(dnName).ShouldNot(BeEmpty(), "Distinguished name of the user is empty")
			Expect(dnName).Should(Equal(os.Getenv("UsernameDN")), "DN name is different than expected: %s", dnName)
		})

		It("Should get DN for user", func() {
			dnName, err := adInstance.GetUserDistinguishedName(os.Getenv("Username"))
			Expect(err).Should(BeNil(), "err is not nil while getting DN for user")
			Expect(dnName).ShouldNot(BeEmpty(), "Distinguished name of the user is empty")
			Expect(dnName).Should(Equal(os.Getenv("UsernameDN")), "DN name is different than expected: %s", dnName)
		})

		It("Should throw error for non existent user", func() {
			dnName, err := adInstance.GetUserDistinguishedName(os.Getenv("UsernameNonexistant"))
			Expect(err).ShouldNot(BeNil(), "err is nil while getting DN for group")
			Expect(dnName).Should(BeEmpty(), "Distinguished name of the group is empty")
		})
	})
})
