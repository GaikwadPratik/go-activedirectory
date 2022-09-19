package goactivedirectory_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	Context("Find user", func() {
		It("Should throw error when username does not exist", func() {
			user, err := adInstance.FindUser(os.Getenv("UsernameNonexistant"))
			Expect(err).ShouldNot(BeNil(), "Error is nil while searching by non existant user")
			Expect(user).Should(BeNil(), "User record received for non existant user: %v", user)
		})

		It("Should throw error when DN does not exist", func() {
			user, err := adInstance.FindUser(os.Getenv("UsernameNonexistantDN"))
			Expect(err).ShouldNot(BeNil(), "Error is nil while searching by non existant DN user")
			Expect(user).Should(BeNil(), "User record received for non existant DN user: %v", user)
		})

		It("Should find the user by username", func() {
			user, err := adInstance.FindUser(os.Getenv("Username"))
			Expect(err).Should(BeNil(), "Error is not nil while searching by existant user: %v", err)
			Expect(user).ShouldNot(BeNil(), "User record is nil existant user")
			Expect(user.DistinguishedName).Should(Equal(os.Getenv("UsernameDN")))
		})

		It("Should find the user by DN", func() {
			user, err := adInstance.FindUser(os.Getenv("UsernameDN"))
			Expect(err).Should(BeNil(), "Error is not nil while searching by DN existant user: %v", err)
			Expect(user).ShouldNot(BeNil(), "User record is nil existant DN user")
			Expect(user.SAMAccountName).Should(Equal(os.Getenv("Username")))
		})
	})

	Context("Find users", func() {
		It("Should get list of users", func() {
			users, err := adInstance.FindUsers()
			Expect(err).Should(BeNil(), "Error is not nil listing users: %v", err)
			Expect(len(users)).Should(BeNumerically(">", 0), "User records are not returned")
		})
	})
})
