package goactivedirectory

import (
	"fmt"
)

// GetGroupDistinguishedName returns the distinguished name for the specified group (cn).
func (ad ActiveDirectory) GetGroupDistinguishedName(groupName string) (string, error) {
	if len(groupName) == 0 {
		return "", fmt.Errorf("groupName is not supplied")
	}
	if isDistinguishedName(groupName) {
		return groupName, nil
	}

	entry, err := ad.searchOne(
		getGroupQueryFilter(groupName),
		[]string{""},
	)
	if err != nil {
		return "", err
	}

	return entry.DN, nil
}

// GetUserDistinguishedName the distinguished name for the specified user (userPrincipalName/email or sAMAccountName).
func (ad ActiveDirectory) GetUserDistinguishedName(userName string) (string, error) {
	if len(userName) == 0 {
		return "", fmt.Errorf("userName is not supplied")
	}
	if isDistinguishedName(userName) {
		return userName, nil
	}

	entry, err := ad.searchOne(
		getUserQueryFilter(userName),
		[]string{""},
	)
	if err != nil {
		return "", err
	}

	return entry.DN, nil
}
