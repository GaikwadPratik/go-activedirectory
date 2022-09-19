package goactivedirectory_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	goactivedirectory "github.com/gaikwadpratik/go-activedirectory"
)

var _ = Describe("MemberOf", func() {
	Context("Group", func() {
		It("Should return error when input is invalid", func() {
			input := goactivedirectory.GroupParentsRequest{
				GroupName: "",
			}
			names, err := adInstance.GetMemberOfForGroup(input)
			Expect(err).ShouldNot(BeNil(), "Error is nil")
			Expect(names).Should(BeEmpty(), "Parent name list is not empty: %v", names)
		})

		It("Should get parent of the group by DN", func() {
			input := goactivedirectory.GroupParentsRequest{
				GroupName: os.Getenv("GroupDNName"),
			}
			names, err := adInstance.GetMemberOfForGroup(input)
			Expect(err).Should(BeNil(), "Error is not nil")
			Expect(names).ShouldNot(BeEmpty(), "Parent name list is empty")
		})

		It("Should get parent of the group by CN", func() {
			input := goactivedirectory.GroupParentsRequest{
				GroupName: os.Getenv("GroupName"),
			}
			names, err := adInstance.GetMemberOfForGroup(input)
			Expect(err).Should(BeNil(), "Error is not nil")
			Expect(names).ShouldNot(BeEmpty(), "Parent name list is empty")
		})

		It("Should return empty parent of the group by DN where non exist", func() {
			input := goactivedirectory.GroupParentsRequest{
				GroupName: os.Getenv("GroupInvalidDNName"),
			}
			names, err := adInstance.GetMemberOfForGroup(input)
			Expect(err).Should(BeNil(), "Error is not nil")
			Expect(names).Should(BeEmpty(), "Parent name list is not empty: %v", names)
		})
	})

	Context("User", func() {
		It("Should return error when input is invalid", func() {
			input := goactivedirectory.UserParentsRequest{
				UserName: "",
			}
			names, err := adInstance.GetMemberOfForUser(input)
			Expect(err).ShouldNot(BeNil(), "Error is nil")
			Expect(names).Should(BeEmpty(), "Parent name list is not empty: %v", names)
		})

		It("Should get parent of the user by DN", func() {
			input := goactivedirectory.UserParentsRequest{
				UserName: os.Getenv("UsernameDN"),
			}
			names, err := adInstance.GetMemberOfForUser(input)
			Expect(err).Should(BeNil(), "Error is not nil")
			Expect(names).ShouldNot(BeEmpty(), "Parent name list is empty")
		})

		It("Should get parent of the user by CN", func() {
			input := goactivedirectory.UserParentsRequest{
				UserName: os.Getenv("Username"),
			}
			names, err := adInstance.GetMemberOfForUser(input)
			Expect(err).Should(BeNil(), "Error is not nil")
			Expect(names).ShouldNot(BeEmpty(), "Parent name list is empty")
		})

		It("Should return empty parent of the group by DN where non exist", func() {
			input := goactivedirectory.UserParentsRequest{
				UserName: os.Getenv("UsernameDNInvalid"),
			}
			names, err := adInstance.GetMemberOfForUser(input)
			Expect(err).Should(BeNil(), "Error is not nil")
			Expect(names).Should(BeEmpty(), "Parent name list is not empty: %v", names)
		})
	})

})
