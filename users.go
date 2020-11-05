package ttbauth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:id`
	Username  string    `json:username`
	Password  string    `json:password`
	Status    int       `json:status`
	CreatedAt time.Time `json:created_at`
}

type UserListEntry struct {
	Id       uuid.UUID `json:id`
	Username string    `json:username`
	Status   int       `json:status`
}

type UserList []UserListEntry

type UserRequest struct {
	Session uuid.UUID `json:session`
	Action  int       `json:id`
	Data    []User
}

type UserResponse struct {
	Status  int    `json:status`
	Message string `json:message`
	Data    User   `json:data`
}

type UserListResponse struct {
	Status  int      `json:status`
	Message string   `json:message`
	Data    UserList `json:data`
}

const (
	StatusNew      = 1
	StatusActive   = 2
	StatusInactive = 3
	StatusDeleted  = 4

	StatusOK               = 0
	StatusInvalidSession   = 1
	StatusPermissionDenied = 2

	ActionList   = 1
	ActionView   = 2
	ActionAdd    = 3
	ActionUpdate = 4
	ActionDelete = 5
)

var (
	ErrInvalidUserStatus = errors.New("Invalid user status")
)

func (user *User) Validate() bool {
	if !user.ValidateUsername() {
		return false
	}

	if !user.ValidateStatus() {
		return false
	}

	return true
}

func (user *User) ValidateUsername() bool {
	// TODO: add validation code here
	return true
}

func (user *User) ValidateStatus() bool {
	return user.Status >= 1 && user.Status <= 4
}

/*
Copyright 2020 Lorenzo Cabrini

Use of this source code is governed by an MIT-style
license that can be found in the LICENSE file or at
https://opensource.org/licenses/MIT.
*/
