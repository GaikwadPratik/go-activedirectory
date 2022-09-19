package goactivedirectory

import (
	"crypto/x509"

	ldap "github.com/go-ldap/ldap/v3"
)

type ActiveDirectory struct {
	conn   *ldap.Conn
	config *ActiveDirectoryConnConfig
}

// ServerConfig is basic configuration required for opening a socket to AD
type ServerConfig struct {
	//Url for connecting to server
	//example may include `ldaps://xyz.lan` or `ldap://xyz.lan`
	Url string
	//Chain of certs required for TLS connection
	RootCAs *x509.CertPool
}

// ActiveDirectoryConnConfig basic configuration used for connecting to AD server
type ActiveDirectoryConnConfig struct {
	*ServerConfig
	//Username to be used to for login. Can be a service account username as well
	//Should always be of the form username@domainname to avoid any confusion
	AdminUsername string
	//Password for the admin or service account
	AdminPassword string
	//BaseDN is the root where the search will happen. If not known, can be found
	// using <cref="GetAttributeOnRootDSE"/>
	BaseDN string
}

type ActiveDirectoryUser struct {
	DistinguishedName string `activedirectory:"distinguishedName"`
	UserPrincipalName string `activedirectory:"userPrincipalName"`
	SAMAccountName    string `activedirectory:"sAMAccountName"`
	SID               string `activedirectory:"objectSid"`
	Mail              string `activedirectory:"mail"`
	// LockoutTime        *time.Time `activedirectory:"lockoutTime"`
	// WhenCreated        *time.Time `activedirectory:"whenCreated"`
	// PwdLastSet         *time.Time `activedirectory:"pwdLastSet"`
	UserAccountControl string `activedirectory:"userAccountControl"`
	EmployeeID         string `activedirectory:"employeeID"`
	Surname            string `activedirectory:"sn"`
	GivenName          string `activedirectory:"givenName"`
	Initials           string `activedirectory:"initials"`
	CommonName         string `activedirectory:"cn"`
	DisplayName        string `activedirectory:"displayName"`
	Comment            string `activedirectory:"comment"`
	Description        string `activedirectory:"description"`
	OU                 string `activedirectory:"ou"`
	ObjectCategory     string `activedirectory:"objectCategory"`
}

func newActiveDirectoryUser() *ActiveDirectoryUser {
	return &ActiveDirectoryUser{}
}

type ActiveDirectoryGroup struct {
	DistinguishedName string   `activedirectory:"distinguishedName"`
	SAMAccountName    string   `activedirectory:"sAMAccountName"`
	CommonName        string   `activedirectory:"cn"`
	Description       string   `activedirectory:"description"`
	SID               string   `activedirectory:"objectSid"`
	ObjectCategory    string   `activedirectory:"objectCategory"`
	Members           []string `activedirectory:"member"`
	//WhenCreated       *time.Time `activedirectory:"whenCreated"`
}

func newActiveDirectoryGroup() *ActiveDirectoryGroup {
	return &ActiveDirectoryGroup{}
}
