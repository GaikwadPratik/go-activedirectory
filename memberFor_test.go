package goactivedirectory_test

import (
	"os"

	goactivedirectory "github.com/gaikwadpratik/go-activedirectory"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MemberFor", func() {
	Context("Group", func() {
		It("Should throw error when group name is not provided", func() {
			request := goactivedirectory.UsersForGroupRequest{
				GroupName: "",
			}
			usrs, err := adInstance.GetUsersForGroup(request)
			Expect(err).ShouldNot(BeNil(), "Error is empty")
			Expect(usrs).Should(BeEmpty(), "List of users is not empty")
		})

		It("Should return empty list of users of the group when non exist", func() {
			request := goactivedirectory.UsersForGroupRequest{
				GroupName: os.Getenv("GroupNonexistantName"),
			}
			usrs, err := adInstance.GetUsersForGroup(request)
			Expect(err).ShouldNot(BeNil(), "Error is empty")
			Expect(usrs).Should(BeEmpty(), "List of users is not empty")
		})

		It("Should get list of user of the group", func() {
			request := goactivedirectory.UsersForGroupRequest{
				GroupName: os.Getenv("GroupName"),
			}
			usrs, err := adInstance.GetUsersForGroup(request)
			Expect(err).Should(BeNil(), "Error is not empty: %v", err)
			Expect(usrs).ShouldNot(BeEmpty(), "List of users is empty")
			Expect(usrs).Should(ContainElement(os.Getenv("UsernameDN")))
		})
	})
})
