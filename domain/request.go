package domain

import "time"

type Auth struct {
	AuthId int
	Username  string
	Email     string
	Password  string
	RoleId    int8
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ClaimedToken struct {
	UserId int
	Username string
	Role int
}