package ttbauth

import (
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
	Status int   `json:status`
	Errors []int `json:errors`
	Data   User  `json:data`
}

type UserListResponse struct {
	Status int      `json:status`
	Data   UserList `json:data`
}

func (user *User) Validate() (bool, []int) {
	var errs []int

	if !user.ValidateUsername() {
		errs = append(errs, ErrInvalidUsername)
	}

	if !user.ValidateStatus() {
		errs = append(errs, ErrInvalidStatus)
	}

	if len(errs) < 1 {
		return true, nil
	} else {
		return false, errs
	}
}

func (user *User) ValidateUsername() bool {
	// TODO: add validation code here
	return true
}

func (user *User) ValidateStatus() bool {
	return user.Status >= StatusNew && user.Status <= StatusDeleted
}

/*
Copyright 2020 Lorenzo Cabrini

Use of this source code is governed by an MIT-style
license that can be found in the LICENSE file or at
https://opensource.org/licenses/MIT.
*/
