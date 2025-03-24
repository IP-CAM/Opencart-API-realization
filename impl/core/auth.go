package core

import (
	"fmt"
	"ocapi/entity"
)

func (c *Core) AuthenticateByToken(token string) (*entity.User, error) {
	if token == "" {
		return nil, fmt.Errorf("token not provided")
	}

	if userName, ok := c.keys[token]; ok {
		return &entity.User{Username: userName}, nil
	}

	userName, err := c.repo.CheckApiKey(token)
	if err == nil {
		c.log.With("username", userName).Debug("User authenticated from database")
		c.keys[token] = userName
		return &entity.User{Username: userName}, nil
	}

	if c.authKey == token {
		userName = "internal"
		c.log.With("username", userName).Debug("User authenticated from config")
		c.keys[token] = userName
		return &entity.User{Username: userName}, nil
	}

	return nil, fmt.Errorf("invalid token")
}
