package goactivedirectory

import (
	"fmt"
	"regexp"
	"strings"

	ldap "github.com/go-ldap/ldap/v3"
)

// parseDistinguishedName Parses the distinguishedName (dn) to remove any invalid characters or to
// properly escape the request.
func parseDistinguishedName(dn string) string {
	if len(dn) == 0 {
		return ""
	}

	// implement escape rules described in https://social.technet.microsoft.com/wiki/contents/articles/5312.active-directory-characters-to-escape.aspx
	temp := strings.Split(dn, ",")
	component := make([]string, 0)
	var sl string
	reg, _ := regexp.Compile("(?i)^(CN|OU|DC)=")
	for _, comp := range temp {
		match := reg.MatchString(comp)
		if len(comp) > 0 && !match {
			// comma was not a component separator but was embedded in a componentvalue e.g. 'CN=Doe\, John'
			sl, component = component[len(component)-1], component[:len(component)-1]
			sl = fmt.Sprintf("%s\\,%s", sl, comp)
			component = append(component, sl)
		} else {
			component = append(component, comp)
		}
	}

	for i := 0; i < len(component); i++ {
		compValue := component[i][3:]
		newValue := ""
		for j := 0; j < len(compValue); j++ {
			char := compValue[j : j+1]
			switch char {
			case "*":
				char = "\\\\2A"
			case "(":
				char = "\\\\28"
			case ")":
				char = "\\\\29"
			case "+":
				char = "\\+"
			case "<":
				char = "\\<"
			case ">":
				char = "\\>"
			case ";":
				char = "\\;"
			case "=":
				char = "\\="
			case " ":
				if j == 0 || j == len(compValue)-1 {
					char = "\\ "
				}
			}
			newValue = fmt.Sprintf("%s%s", newValue, char)
		}
		component[i] = fmt.Sprintf("%s%s", component[i][0:3], newValue)
	}
	return strings.Join(component, ",")
}

// isDistinguishedName Checks to see if the value is a distinguished name.
func isDistinguishedName(value string) bool {
	if len(value) == 0 {
		return false
	}
	reg, _ := regexp.Compile("(?i)(([^=]+=.+),?)+")
	return reg.Match([]byte(value))
}

// getUserQueryFilter Gets the ActiveDirectory LDAP query string for a user search.
func getUserQueryFilter(username string) string {
	if len(username) == 0 {
		return "(objectCategory=User)"
	}
	if isDistinguishedName(username) {
		return fmt.Sprintf("(&(objectCategory=User)(distinguishedName=%s))", parseDistinguishedName(username))
	}

	return fmt.Sprintf("(&(objectCategory=User)(|(sAMAccountName=%s)(userPrincipalName=%s)))", username, username)
}

// getGroupQueryFilter Gets the ActiveDirectory LDAP query string for a group search.
func getGroupQueryFilter(groupName string) string {
	if len(groupName) == 0 {
		return "(objectCategory=Group)"
	}
	if isDistinguishedName(groupName) {
		return fmt.Sprintf("(&(objectCategory=Group)(distinguishedName=%s))", parseDistinguishedName(groupName))
	}
	return fmt.Sprintf("(&(objectCategory=Group)(cn=%s))", groupName)
}

func isGroupResult(item *ldap.Entry) bool {
	if item == nil {
		return false
	}

	if len(item.Attributes) == 0 {
		return false
	}

	for _, attr := range item.Attributes {
		if attr.Name == "groupType" {
			return true
		}
	}

	return false
}
