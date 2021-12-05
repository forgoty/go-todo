package models

import "time"

type SignedInUser struct {
	UserId      int64
	Login       string
	Name        string
	Email       string
	ApiKeyId    int64
	OrgCount    int
	IsAnonymous bool
	LastSeenAt  time.Time
}
