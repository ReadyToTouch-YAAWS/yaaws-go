// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package dbs

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type SocialProvider string

const (
	SocialProviderGithub    SocialProvider = "github"
	SocialProviderGitlab    SocialProvider = "gitlab"
	SocialProviderBitbucket SocialProvider = "bitbucket"
	SocialProviderFigma     SocialProvider = "figma"
	SocialProviderDribbble  SocialProvider = "dribbble"
	SocialProviderBehance   SocialProvider = "behance"
	SocialProviderLinkedin  SocialProvider = "linkedin"
	SocialProviderXing      SocialProvider = "xing"
	SocialProviderGoogle    SocialProvider = "google"
)

func (e *SocialProvider) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = SocialProvider(s)
	case string:
		*e = SocialProvider(s)
	default:
		return fmt.Errorf("unsupported scan type for SocialProvider: %T", src)
	}
	return nil
}

type NullSocialProvider struct {
	SocialProvider SocialProvider
	Valid          bool // Valid is true if SocialProvider is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullSocialProvider) Scan(value interface{}) error {
	if value == nil {
		ns.SocialProvider, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.SocialProvider.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullSocialProvider) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.SocialProvider), nil
}

type User struct {
	ID        int64
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
}

type UserOnlineDailyCountStat struct {
	CreatedAt pgtype.Date
	UserCount int64
}

type UserOnlineDailyStat struct {
	CreatedAt pgtype.Date
	UserID    int64
}

type UserOnlineHourlyStat struct {
	UserID    int64
	CreatedAt pgtype.Timestamp
}

type UserSocialProfile struct {
	ID                   int64
	UserID               int64
	SocialProvider       SocialProvider
	SocialProviderUserID string
	Email                string
	Username             string
	Name                 string
	CreatedAt            pgtype.Timestamp
	UpdatedAt            pgtype.Timestamp
	DeletedAt            pgtype.Timestamp
}

type UserSocialProfileChangeHistory struct {
	ID                  int64
	UserID              int64
	UserSocialProfileID int64
	Email               string
	Username            string
	CreatedAt           pgtype.Timestamp
	DeletedAt           pgtype.Timestamp
}
