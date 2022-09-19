package goactivedirectory

import (
	"fmt"
)

// FindUser Retrieves the specified user.
// username The username to retrieve information about.
// Optionally can pass in the distinguishedName (dn) of the user to retrieve.
func (ad ActiveDirectory) FindUser(username string) (*ActiveDirectoryUser, error) {
	if len(username) == 0 {
		return nil, fmt.Errorf("input username is not valid")
	}
	entry, err := ad.searchOne(
		getUserQueryFilter(username),
		defaultUserAttributes,
	)

	if err != nil {
		return nil, err
	}

	usr := newActiveDirectoryUser()

	err = fillObjFromAttrMap(usr, entry.Attributes)

	if err != nil {
		return nil, err
	}

	return usr, nil
}

// FindUsers lists all the users
func (ad ActiveDirectory) FindUsers() ([]*ActiveDirectoryUser, error) {
	filter := "(&(|(objectClass=user)(objectClass=person))(!(objectClass=computer))(!(objectClass=group)))"
	entries, err := ad.search(
		filter,
		defaultUserAttributes,
		1000,
	)

	if err != nil {
		return nil, err
	}

	rtnVal := make([]*ActiveDirectoryUser, 0)

	for _, entry := range entries {
		usr := newActiveDirectoryUser()
		err = fillObjFromAttrMap(usr, entry.Attributes)
		if err != nil {
			return nil, err
		}
		rtnVal = append(rtnVal, usr)
	}

	return rtnVal, nil
}
