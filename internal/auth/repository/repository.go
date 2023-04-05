package repository

import (
	"context"

	"miniWiki/internal/auth/model"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r AuthRepository) GetSession(ctx context.Context, sessionID string) (*model.Session, error) {
	var session model.Session
	err := r.db.WithContext(ctx).Take(&session, "session_id = ?", sessionID).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r AuthRepository) DeleteSession(ctx context.Context, sessionID string) error {
	session := model.Session{
		SessionID: sessionID,
	}
	err := r.db.WithContext(ctx).Delete(&session).Error
	return err
}

func (r AuthRepository) UpdateSession(ctx context.Context, sessionID string, session model.Session) error {
	err := r.db.WithContext(ctx).Model(model.Session{}).Where("session_id = ?", sessionID).Updates(session).Error
	return err
}

func (r AuthRepository) CreateSession(ctx context.Context, session model.Session) error {
	err := r.db.WithContext(ctx).Create(&session).Error
	return err
}
