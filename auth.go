package ttbauth

import (
	"github.com/google/uuid"
)

type AuthRequest struct {
	Username string `json:username`
	Password string `json:password`
}

type LogoutRequest struct {
	Session uuid.UUID `json:session`
}

type AuthResponse struct {
	Status  int       `json:status`
	Message string    `json:message`
	Session uuid.UUID `json:session`
}

const (
	StatusAuthOK       = 0
	StatusAuthFailed   = 1
	StatusAuthInactive = 2


/*
Copyright 2020 Lorenzo Cabrini

Use of this source code is governed by an MIT-style
license that can be found in the LICENSE file or at
https://opensource.org/licenses/MIT.
*/
