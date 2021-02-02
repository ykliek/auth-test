package model

import (
	"github.com/twinj/uuid"
	"api-auth-test/auth"
)

type Auth struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID   uint64 `gorm:";not null;" json:"user_id"`
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
}

func (s *Server) FetchAuth(authD *auth.Details) (*Auth, error) {
	au := &Auth{}
	err := s.DB.Debug().Where("user_id AND auth_uuid = ?", authD.UserId, authD.AuthUid).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}

func (s *Server) DeleteAuth(authD *auth.Details) error {
	au := &Auth{}
	db := s.DB.Debug().Where("where user_id = ? AND auth_uuid = ?", authD.UserId, authD.AuthUid).Take(&au).Delete(&au)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (s *Server) CreateAuth(userId uint64) (*Auth, error) {
	au := &Auth{}
	au.AuthUUID = uuid.NewV4().String()
	au.UserID = userId
	err := s.DB.Debug().Create(&au).Error
	if err != nil {
		return nil, err
	}
	return au, err
}


