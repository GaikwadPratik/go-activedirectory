package goactivedirectory

import "fmt"

// Authenticate Authenticates the username and password by doing a simple bind with the specified credentials.
// username may be either the sAMAccountName or the userPrincipalName.
func (ad ActiveDirectory) Authenticate(username string, password string) (bool, error) {
	if len(username) == 0 {
		return false, fmt.Errorf("username is not provided")
	}

	if len(password) == 0 {
		return false, fmt.Errorf("password is not provided")
	}

	upn, err := ad.GetUPN(username)
	if err != nil {
		return false, err
	}

	conn, err := dialConnection(ad.config.ServerConfig)

	if err != nil {
		return false, err
	}

	err = conn.Bind(upn, password)
	if err != nil {
		return false, err
	}

	err = conn.Unbind()
	if err != nil {
		return false, err
	}

	return true, nil
}
