package goactivedirectory_test

import (
	"log"
	"os"
	"testing"

	goactivedirectory "github.com/gaikwadpratik/go-activedirectory"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var adInstance *goactivedirectory.ActiveDirectory

func TestGoActivedirectory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoActivedirectory Suite")
}

var _ = BeforeSuite(func() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Some basic validations
	Expect(os.Getenv("AdminUsername")).ShouldNot(BeEmpty(), "Please make sure AdminUsername is set correctly.")
	Expect(os.Getenv("AdminPassword")).ShouldNot(BeEmpty(), "Please make sure AdminPassword is set correctly.")
	Expect(os.Getenv("ldapUrl")).ShouldNot(BeEmpty(), "Please make sure ldapUrl is set correctly.")

	conf := goactivedirectory.ActiveDirectoryConnConfig{
		ServerConfig: &goactivedirectory.ServerConfig{
			Url: os.Getenv("ldapUrl"),
		},
		AdminUsername: os.Getenv("AdminUsername"),
		AdminPassword: os.Getenv("AdminPassword"),
		BaseDN:        os.Getenv("BaseDN"),
	}

	adInstance, err = goactivedirectory.NewActiveDirectory(&conf)
	if err != nil {
		log.Fatal(err)
	}
})

var _ = AfterSuite(func() {
	if adInstance == nil {
		return
	}
	err := adInstance.Cleanup()
	if err != nil {
		log.Fatal(err)
	}
})
