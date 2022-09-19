package goactivedirectory_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Group", func() {
	Context("Find group", func() {
		It("Should throw error when group does not exist", func() {
			group, err := adInstance.FindGroup(os.Getenv("GroupNonexistantName"))
			Expect(err).ShouldNot(BeNil(), "Error is nil while searching by non existant group")
			Expect(group).Should(BeNil(), "User record received for non existant user: %v", group)
		})

		It("Should throw error when DN does not exist", func() {
			group, err := adInstance.FindGroup(os.Getenv("GroupNonexistantDNName"))
			Expect(err).ShouldNot(BeNil(), "Error is nil while searching by non existant DN user")
			Expect(group).Should(BeNil(), "User record received for non existant DN user: %v", group)
		})

		It("Should find the group by groupname", func() {
			group, err := adInstance.FindGroup(os.Getenv("GroupName"))
			Expect(err).Should(BeNil(), "Error is not nil while searching by existant user: %v", err)
			Expect(group).ShouldNot(BeNil(), "User record is nil existant user")
			Expect(group.DistinguishedName).Should(Equal(os.Getenv("GroupDNName")))
		})

		It("Should find the group by DN", func() {
			group, err := adInstance.FindGroup(os.Getenv("GroupDNName"))
			Expect(err).Should(BeNil(), "Error is not nil while searching by DN existant user: %v", err)
			Expect(group).ShouldNot(BeNil(), "User record is nil existant DN user")
			Expect(group.SAMAccountName).Should(Equal(os.Getenv("GroupName")))
		})
	})

	Context("Find groups", func() {
		It("Should get list of groups", func() {
			groups, err := adInstance.FindGroups()
			Expect(err).Should(BeNil(), "Error is not nil listing groups: %v", err)
			Expect(len(groups)).Should(BeNumerically(">", 0), "Group records are not returned")
		})
	})
})
