package goactivedirectory_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Authenticate", func() {
	Context("User", func() {
		It("Should throw error when username is not provided", func() {
			result, err := adInstance.Authenticate("", "")
			Expect(err).ShouldNot(BeNil(), "Username validation is empty")
			Expect(result).Should(Equal(false), "result is true")
		})

		It("Should throw error when password is not provided", func() {
			result, err := adInstance.Authenticate(os.Getenv("Username"), "")
			Expect(err).ShouldNot(BeNil(), "Password validation is empty")
			Expect(result).Should(Equal(false), "result is true")
		})

		It("Should throw error when credentials(Username) are wrong", func() {
			result, err := adInstance.Authenticate(os.Getenv("UsernameNonexistant"), os.Getenv("Password"))
			Expect(err).ShouldNot(BeNil(), "Invalid credential error is empty")
			Expect(result).Should(Equal(false), "result is true")
		})

		It("Should throw error when credentials(Password) are wrong", func() {
			result, err := adInstance.Authenticate(os.Getenv("Username"), os.Getenv("PasswordInvalid"))
			Expect(err).ShouldNot(BeNil(), "Invalid credential error is empty")
			Expect(result).Should(Equal(false), "result is true")
		})

		It("Should validate credentials(sAM)", func() {
			result, err := adInstance.Authenticate(os.Getenv("Username"), os.Getenv("Password"))
			Expect(err).Should(BeNil(), "Error is not empty: %v", err)
			Expect(result).Should(Equal(true), "result is false")
		})

		It("Should validate credentials(UPN)", func() {
			result, err := adInstance.Authenticate(os.Getenv("Upn"), os.Getenv("Password"))
			Expect(err).Should(BeNil(), "Error is not empty: %v", err)
			Expect(result).Should(Equal(true), "result is false")
		})
	})
})
