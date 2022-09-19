package goactivedirectory

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

// Returns root configuration of the server
func GetRootDSE(config *ServerConfig) ([]*ldap.Entry, error) {
	searchRequest := ldap.NewSearchRequest(
		"",
		ldap.ScopeBaseObject,
		ldap.NeverDerefAliases,
		1000,
		0,
		false,
		"(&(supportedLDAPVersion=*)(objectClass=*))",
		[]string{"*"},
		nil,
	)

	conn, err := dialConnection(config)

	if err != nil {
		return nil, err
	}

	searchResult, err := conn.Search(searchRequest)
	if err != nil {
		return nil, err
	}
	return searchResult.Entries, nil
}

// GetAttributeOnRootDSE returns a specific attribute value
func GetAttributeOnRootDSE(config *ServerConfig, attributeName string) (string, error) {

	entries, err := GetRootDSE(config)

	if err != nil {
		return "", err
	}

	attributeValue := ""

	for _, entry := range entries {
		attributeValue = entry.GetAttributeValue(attributeName)
		if len(attributeValue) > 0 {
			break
		}
	}

	if len(attributeValue) == 0 {
		return "", fmt.Errorf("%s not returned by AD", attributeName)
	}

	return attributeValue, nil
}
