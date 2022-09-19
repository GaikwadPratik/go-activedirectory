package goactivedirectory

import (
	"crypto/tls"
	"fmt"
	"net/mail"
	"strings"

	ldap "github.com/go-ldap/ldap/v3"
)

func dialConnection(config *ServerConfig) (*ldap.Conn, error) {
	var conn *ldap.Conn
	var err error
	if strings.HasPrefix(config.Url, "ldaps") {
		tlsConfig := &tls.Config{InsecureSkipVerify: true}

		if config.RootCAs != nil {
			tlsConfig.RootCAs = config.RootCAs
			tlsConfig.InsecureSkipVerify = false
		}

		conn, err = ldap.DialTLS("tcp", config.Url, tlsConfig)
		if err != nil {
			return nil, fmt.Errorf("connection error: %w", err)
		}
	} else {
		conn, err = ldap.DialURL(config.Url)
		if err != nil {
			return nil, fmt.Errorf("connection error: %w", err)
		}
	}
	return conn, nil
}

// SearchOne returns the single entry for the given search criteria or an error if one occurred.
// An error is returned if exactly one entry is not returned.
func (ad ActiveDirectory) searchOne(filter string, attrs []string) (*ldap.Entry, error) {
	search := ldap.NewSearchRequest(
		ad.config.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		1,
		0,
		false,
		filter,
		attrs,
		nil,
	)

	result, err := ad.conn.Search(search)
	if err != nil {
		if e, ok := err.(*ldap.Error); ok {
			if e.ResultCode == ldap.LDAPResultSizeLimitExceeded {
				return nil, fmt.Errorf(`search error "%s": more than one entries returned`, filter)
			}
		}

		return nil, fmt.Errorf(`search error "%s": %w`, filter, err)
	}

	if len(result.Entries) == 0 {
		return nil, fmt.Errorf(`search error "%s": no entries returned`, filter)
	}

	return result.Entries[0], nil
}

func (ad ActiveDirectory) search(filter string, attrs []string, sizeLimit int) ([]*ldap.Entry, error) {
	search := ldap.NewSearchRequest(
		ad.config.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		sizeLimit,
		0,
		false,
		filter,
		attrs,
		nil,
	)
	result, err := ad.conn.Search(search)
	if err != nil {
		return nil, fmt.Errorf(`search error "%s": %w`, filter, err)
	}

	return result.Entries, nil
}

var defaultUserAttributes = []string{
	"distinguishedName", "userPrincipalName", "sAMAccountName", "objectSID", "mail",
	"userAccountControl", "employeeID", "sn", "givenName", "initials", "cn",
	"displayName", "comment", "description", "ou", "objectCategory",
	//"lockoutTime", "whenCreated", "pwdLastSet",
}

var defaultGroupAttributes = []string{
	"cn", "description", "distinguishedName", "objectCategory", "sAMAccountName", "objectCategory",
	//"whenCreated",
}

// getDomain returns the domain derived from BaseDN or an error if misconfigured.
func (ad ActiveDirectory) getDomain() (string, error) {
	domain := ""
	for _, v := range strings.Split(strings.ToLower(ad.config.BaseDN), ",") {
		if trimmed := strings.TrimSpace(v); strings.HasPrefix(trimmed, "dc=") {
			domain = domain + "." + trimmed[3:]
		}
	}
	if len(domain) <= 1 {
		return "", fmt.Errorf("invalid BaseDN")
	}
	return domain[1:], nil
}

// GetUPN returns the userPrincipalName for the given username or an error if misconfigured.
func (ad ActiveDirectory) GetUPN(username string) (string, error) {
	if _, err := mail.ParseAddress(username); err == nil {
		return username, nil
	}

	domain, err := ad.getDomain()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s@%s", username, domain), nil
}
