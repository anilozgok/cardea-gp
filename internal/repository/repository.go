package repository

import (
	"context"
	"github.com/anilozgok/cardea-gp/internal/entities"
	"gorm.io/gorm"
)

type Repository interface {
	CreateNewUser(ctx context.Context, user *entities.User) error
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
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
	return user, tx.Error
}
