package models

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserID int    `json:"userID"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func (u *User) Validate() error {
	if err := validateUsername(u.Username); err != nil {
		return err
	}
	if err := validatePassword(u.Password); err != nil {
		return err
	}
	if err := validateRole(u.Role); err != nil {
		return err
	}
	return nil
}

func validateUsername(username string) error {
	if strings.TrimSpace(username) == "" {
		return errors.New("username must not be empty")
	}
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	if len(username) > 50 {
		return errors.New("username must be less than 50 characters long")
	}
	return nil
}

func (c *Credentials) Validate() error {
	if err := validateUsername(c.Username); err != nil {
		return err
	}
	if err := validatePassword(c.Password); err != nil {
		return err
	}
	return nil
}
func validatePassword(password string) error {
	if strings.TrimSpace(password) == "" {
		return errors.New("password must not be empty")
	}
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	return nil
}

func validateRole(role string) error {
	if role != "user" && role != "admin" {
		return errors.New("role must be either 'user' or 'admin'")
	}
	return nil
}


