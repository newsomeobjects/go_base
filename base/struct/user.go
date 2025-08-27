package main

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	// 所有字段都是私有的
	id        int
	username  string
	email     string
	createdAt time.Time
	isActive  bool
}

// NewUser 公开的构造函数
func NewUser(username, email string) *User {
	return &User{
		username:  username,
		email:     email,
		createdAt: time.Now(),
		isActive:  true,
	}
}

// ID 公开的getter方法
func (u *User) ID() int {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Email() string {
	return u.email
}

func (u *User) IsActive() bool {
	return u.isActive
}

// SetEmail 有控制的setter方法
func (u *User) SetEmail(newEmail string) error {
	if !isValidEmail(newEmail) {
		return errors.New("invalid email format")
	}
	u.email = newEmail
	return nil
}

// Deactivate 业务方法
func (u *User) Deactivate() {
	u.isActive = false
}

// 私有辅助方法
func isValidEmail(email string) bool {
	return strings.Contains(email, "@")
}
