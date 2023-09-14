package table

import (
	"mid/core/auth"
	customerrors "mid/core/errors"
	"mid/core/var/common"
	"time"

	"gorm.io/gorm"
)

const (
	StatusSuccess = "SUCCESS"
	StatusFailed  = "FAILED"
	StatusPending = "PENDING"
)

func (SmsQueue) TableName() string {
	return "public.table_SmsQueue"
}

type SmsQueue struct {
	Id          int `gorm:"primaryKey;"`
	UserId      int
	PhoneNumber string
	SmsBody     string
	TryCount    int
	Status      string

	CreatedDate *time.Time
	IsDeleted   bool
}

func (u *SmsQueue) BeforeCreate(tx *gorm.DB) (err error) {

	now := time.Now()
	user, ok := tx.Statement.Context.Value(common.LabelUser).(*auth.User)

	if !ok {
		return &customerrors.NoContextFound{}
	}

	if user == nil {
		return &customerrors.NoContextFound{}
	}

	u.UserId = user.UserID
	u.CreatedDate = &now
	u.Status = StatusPending

	return nil

}

func (u *SmsQueue) BeforeUpdate(tx *gorm.DB) (err error) {

	user, ok := tx.Statement.Context.Value(common.LabelUser).(*auth.User)

	if !ok {
		return &customerrors.NoContextFound{}
	}

	if user == nil {
		return &customerrors.NoContextFound{}
	}

	return nil

}
