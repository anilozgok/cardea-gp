package database

import (
	"context"
	"errors"
	"github.com/anilozgok/cardea-gp/internal/model/entities"
	"gorm.io/gorm"
)

type Repository interface {
	CreateNewUser(ctx context.Context, user *entities.User) error
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUsers(ctx context.Context) ([]entities.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateNewUser(ctx context.Context, user *entities.User) error {
	tx := r.db.WithContext(ctx).Create(user)
	return tx.Error
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user := new(entities.User)
	tx := r.db.WithContext(ctx).Where("email = ?", email).First(user)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, tx.Error
}

func (r *repository) GetUsers(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	tx := r.db.WithContext(ctx).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}
