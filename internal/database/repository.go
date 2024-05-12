package database

import (
	"context"
	"errors"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"gorm.io/gorm"
)

type Repository interface {
	CreateNewUser(ctx context.Context, user *entity.User) error
	CreateWorkout(ctx context.Context, workout *entity.Workout) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserById(ctx context.Context, id uint) (*entity.User, error)
	GetUsers(ctx context.Context) ([]entity.User, error)
	ListWorkoutByUserId(ctx context.Context, userId uint) ([]entity.Workout, error)
	ListWorkoutByCoachId(ctx context.Context, coachId uint) ([]entity.Workout, error)
	UpdatePassword(ctx context.Context, password string, user entity.User) error
	// profile related methods
	CreateProfile(ctx context.Context, profile *entity.Profile) error
	GetProfileByUserId(ctx context.Context, userId uint) (*entity.Profile, error)
	UpdateProfile(ctx context.Context, profile *entity.Profile) error
	AddPhoto(ctx context.Context, userId uint, photoURL string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateNewUser(ctx context.Context, user *entity.User) error {
	tx := r.db.WithContext(ctx).Create(user)
	return tx.Error
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	tx := r.db.WithContext(ctx).Where("email = ?", email).First(user)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, tx.Error
}

func (r *repository) GetUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	tx := r.db.WithContext(ctx).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (r *repository) GetUserById(ctx context.Context, id uint) (*entity.User, error) {
	user := new(entity.User)
	tx := r.db.WithContext(ctx).Where(&entity.User{Model: gorm.Model{ID: id}}).First(&user)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, tx.Error
}

func (r *repository) CreateWorkout(ctx context.Context, workout *entity.Workout) error {
	tx := r.db.WithContext(ctx).Create(workout)
	return tx.Error
}

func (r *repository) ListWorkoutByUserId(ctx context.Context, userId uint) ([]entity.Workout, error) {
	var workouts []entity.Workout
	tx := r.db.WithContext(ctx).Where(&entity.Workout{UserId: userId}).Find(&workouts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return workouts, nil
}

func (r *repository) ListWorkoutByCoachId(ctx context.Context, coachId uint) ([]entity.Workout, error) {
	var workouts []entity.Workout
	tx := r.db.WithContext(ctx).Where(&entity.Workout{CoachId: coachId}).Find(&workouts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return workouts, nil
}

func (r *repository) UpdatePassword(ctx context.Context, password string, user entity.User) error {
	user.Password = password
	tx := r.db.WithContext(ctx).Save(user)

	return tx.Error
}

// Profile related methods

func (r *repository) CreateProfile(ctx context.Context, profile *entity.Profile) error {
	tx := r.db.WithContext(ctx).Create(profile)
	return tx.Error
}

func (r *repository) GetProfileByUserId(ctx context.Context, userId uint) (*entity.Profile, error) {
	profile := new(entity.Profile)
	tx := r.db.WithContext(ctx).Where(&entity.Profile{UserId: userId}).First(profile)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return profile, tx.Error
}

func (r *repository) UpdateProfile(ctx context.Context, profile *entity.Profile) error {
	tx := r.db.WithContext(ctx).Save(profile)
	return tx.Error
}

func (r *repository) AddPhoto(ctx context.Context, userId uint, photoURL string) error {
	profile, err := r.GetProfileByUserId(ctx, userId)
	if err != nil {
		return err
	}

	if profile == nil {
		return errors.New("profile not found")
	}

	profile.Photos = append(profile.Photos, photoURL)

	return r.UpdateProfile(ctx, profile)
}
