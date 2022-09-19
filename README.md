# go-activedirectory
go-activedirectory is an ldap client around [go-ldap](https://github.com/go-ldap/ldap) for authN (authentication) and authZ (authorization) for Microsoft Active Directory with range retrieval support for large Active Directory installations. 

## Installing

Using Go modules
```
go get github.com/gaikwadpratik/go-activedirectory
```

**Dependencies:**
* [github.com/go-ldap/ldap](https://github.com/go-ldap/ldap)

# Usage

**Example:**
Basic usage
```go
conf := goactivedirectory.ActiveDirectoryConnConfig{
    ServerConfig: &goactivedirectory.ServerConfig{
		Url: os.Getenv("ldapUrl"),
	},
	AdminUsername: os.Getenv("AdminUsername"),
	AdminPassword: os.Getenv("AdminPassword"),
	BaseDN:        os.Getenv("DefaultNamingContext"),
}

adInstance, err = goactivedirectory.NewActiveDirectory(&conf)
if err != nil {
	log.Fatal(err)
}
```
`adInstance` now can be used to perform all other functionality from the library.

if `BaseDN` is not known, it can be retrieved with below example:
```go
serverConfig := &activedirectory.ServerConfig{
	Url: os.Getenv("ldapUrl"),
}
val, err := activedirectory.GetAttributeOnRootDSE(serverConfig, "defaultNamingContext")
```

The `username` and `password` specified in the configuration are what are used for user and group lookup operations. So they should be of an _elevated_ or _admin_ user

# Cloning and Testing
Clone the repo using 
```
git clone git@github.com:GaikwadPratik/go-activedirectory.git
```

Run test cases run `ginkgo -r` after setting below Environment variables in `.env` file
| Name                    | Description |
| ----------------------- | ------------- |
| ldapUrl                 | Url of server to connect (`ldaps://example.com` or `ldap://example.com`) |
| BaseDN                  | LDAP Base DN - for testing the root DN is recommended, e.g. `DC=example,DC=com` |
| AdminUsername           | userPrincipalName (user@domain.tld) of admin user |
| AdminPassword           | Password of admin user |
| GroupName               | commonName of a test group that DOES exist |
| GroupNonexistantName    | commonName of a test group that does NOT exist  |
| GroupDNName             | distinguishedName of a test group DOES exist |
| GroupInvalidDNName      | distinguishedName of a test group that does NOT exist or is wrong  |
| GroupNonexistantDNName  | distinguishedName of a test group that does NOT exist  |
| Username                | commonName of a test user that DOES exist |
| UsernameNonexistant     | commonName of a test user that does NOT exist |
| UsernameDN              | distinguishedName of a test user DOES exist |
| UsernameDNInvalid       | distinguishedName of a test user that does NOT exist or is wrong  |
| UsernameNonexistantDN   | distinguishedName of a test user that does NOT exist  |
| Password                | Password of test user |
| PasswordInvalid         | Invalid or some random password |
| Upn                     | userPrincipalName of a test user that will be used |

