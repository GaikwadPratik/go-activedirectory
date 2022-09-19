package goactivedirectory

import (
	"fmt"

	ldap "github.com/go-ldap/ldap/v3"
)

// GroupParentsRequest represents request for getting parents of a group by DN
type GroupParentsRequest struct {
	GroupName string
	//TODO: FIXME: add recursive code
	//Recurssive             bool
}

// GetMemberOfForGroup For the specified group, get all of the groups that the group is a member of.
func (ad ActiveDirectory) GetMemberOfForGroup(input GroupParentsRequest) ([]string, error) {
	if len(input.GroupName) == 0 {
		return nil, fmt.Errorf("input groupname must be specified")
	}
	if !isDistinguishedName(input.GroupName) {
		groupDNName, err := ad.GetGroupDistinguishedName(input.GroupName)
		if err != nil {
			return nil, err
		}
		input.GroupName = groupDNName
	}
	searchResult, err := ad.search(
		fmt.Sprintf("(member=%s)", ldap.EscapeFilter(parseDistinguishedName(input.GroupName))),
		[]string{"groupType"},
		1000,
	)
	if err != nil {
		return nil, err
	}

	rtnVal := make([]string, 0)

	for _, entry := range searchResult {
		rtnVal = append(rtnVal, entry.DN)
	}

	return rtnVal, nil
}

// UserParentsRequest represents request for getting parents of a user by DN
type UserParentsRequest struct {
	UserName string
	//TODO: FIXME: add recursive code
	//Recurssive             bool
}

// GetMemberOfForUser For the specified username, get all of the groups that the user is a member of
// returns the DN of the groups
func (ad ActiveDirectory) GetMemberOfForUser(input UserParentsRequest) ([]string, error) {
	if len(input.UserName) == 0 {
		return nil, fmt.Errorf("input username must be specified")
	}
	if !isDistinguishedName(input.UserName) {
		userDNName, err := ad.GetUserDistinguishedName(input.UserName)
		if err != nil {
			return nil, err
		}
		input.UserName = userDNName
	}
	searchResult, err := ad.search(
		fmt.Sprintf("(member=%s)", ldap.EscapeFilter(parseDistinguishedName(input.UserName))),
		[]string{"groupType"},
		1000,
	)
	if err != nil {
		return nil, err
	}

	rtnVal := make([]string, 0)

	for _, entry := range searchResult {
		rtnVal = append(rtnVal, entry.DN)
	}

	return rtnVal, nil
}
