package goactivedirectory

import (
	"fmt"
)

// NewActiveDirectory initiates a new connection based on provided configuration
func NewActiveDirectory(config *ActiveDirectoryConnConfig) (*ActiveDirectory, error) {
	if config == nil {
		return nil, fmt.Errorf("configuration is not provided")
	}

	if config.ServerConfig == nil {
		return nil, fmt.Errorf("server config is missing")
	}

	if len(config.ServerConfig.Url) == 0 {
		return nil, fmt.Errorf("server config url is missing")
	}

	if len(config.AdminUsername) == 0 {
		return nil, fmt.Errorf("missing admin username")
	}

	if len(config.AdminUsername) == 0 {
		return nil, fmt.Errorf("missing admin password")
	}

	if len(config.BaseDN) == 0 {
		return nil, fmt.Errorf("missing admin password")
	}

	conn, err := dialConnection(config.ServerConfig)

	if err != nil {
		return nil, err
	}

	err = conn.Bind(config.AdminUsername, config.AdminPassword)
	if err != nil {
		return nil, err
	}

	return &ActiveDirectory{
		conn:   conn,
		config: config,
	}, nil
}

// Cleanup unbinds a user and closes any open connection
func (ad ActiveDirectory) Cleanup() error {
	if ad.conn == nil {
		return nil
	}
	err := ad.conn.Unbind()
	if err != nil {
		return err
	}

	return nil
}
