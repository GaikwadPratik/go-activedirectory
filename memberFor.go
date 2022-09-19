package goactivedirectory

import (
	"fmt"
)

type UsersForGroupRequest struct {
	GroupName string
	//TODO: FIXME: add recursive code
	//Recurssive             bool
}

// GetUsersForGroup For the specified group, retrieve all of the users that belong to the group.
// returns DN of the users
func (ad ActiveDirectory) GetUsersForGroup(request UsersForGroupRequest) ([]string, error) {
	if len(request.GroupName) == 0 {
		return nil, fmt.Errorf("groupName is missing")
	}

	groupAttributes := []string{"member"}
	grp, err := ad.findGroup(request.GroupName, groupAttributes)
	if err != nil {
		return nil, err
	}

	members := grp.GetAttributeValues("member")

	if len(members) == 0 {
		return members, nil
	}

	filter := ""

	for _, mem := range members {
		filter += fmt.Sprintf("(distinguishedName=%s)", parseDistinguishedName(mem))
	}

	filter = fmt.Sprintf("(&(|(objectCategory=User)(objectCategory=Group))(|%s))", filter)

	searchResult, err := ad.search(
		filter,
		append(defaultUserAttributes, "groupType"),
		1000,
	)

	if err != nil {
		return nil, err
	}

	rtnVal := make([]string, 0)

	for _, entry := range searchResult {
		if !isGroupResult(entry) {
			rtnVal = append(rtnVal, entry.DN)
			continue
		}
		//TODO: Recurssive code for getting users of children group
	}

	return rtnVal, nil
}
