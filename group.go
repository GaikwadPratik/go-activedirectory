package goactivedirectory

import (
	"fmt"

	ldap "github.com/go-ldap/ldap/v3"
)

func (ad ActiveDirectory) findGroup(groupName string, groupAttributs []string) (*ldap.Entry, error) {
	entries, err := ad.search(
		getGroupQueryFilter(groupName),
		groupAttributs,
		1000,
	)

	if err != nil {
		return nil, err
	}

	if len(entries) == 0 {
		return nil, fmt.Errorf("no group found matching for groupname: %s", groupName)
	}

	return entries[0], nil
}

// FindGroup Retrieves the specified group.
// groupname can be CommonName(cn) or distinguishedName (dn).
func (ad ActiveDirectory) FindGroup(groupName string) (*ActiveDirectoryGroup, error) {
	if len(groupName) == 0 {
		return nil, fmt.Errorf("groupname is not provided")
	}

	entry, err := ad.findGroup(groupName, defaultGroupAttributes)
	if err != nil {
		return nil, err
	}

	grp := newActiveDirectoryGroup()

	err = fillObjFromAttrMap(grp, entry.Attributes)

	if err != nil {
		return nil, err
	}

	return grp, err
}

// FindGroups lists all the groups
func (ad ActiveDirectory) FindGroups() ([]*ActiveDirectoryGroup, error) {
	filter := "(&(objectClass=group)(!(objectClass=computer))(!(objectClass=user))(!(objectClass=person)))"
	entries, err := ad.search(
		filter,
		defaultGroupAttributes,
		1000,
	)

	if err != nil {
		return nil, err
	}

	rtnVal := make([]*ActiveDirectoryGroup, 0)

	for _, entry := range entries {
		grp := newActiveDirectoryGroup()
		err = fillObjFromAttrMap(grp, entry.Attributes)
		if err != nil {
			return nil, err
		}
		rtnVal = append(rtnVal, grp)
	}

	return rtnVal, nil
}
