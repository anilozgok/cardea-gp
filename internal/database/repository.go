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
	CreateProfile(ctx context.Context, profile *entity.Profile) error
	GetProfileByUserId(ctx context.Context, userId uint) (*entity.Profile, error)
	UpdateProfile(ctx context.Context, profile *entity.Profile) error
	AddPhoto(ctx context.Context, photo *entity.Image) error
	GetWorkoutById(ctx context.Context, id uint) (*entity.Workout, error)
	UpdateWorkout(ctx context.Context, workout entity.Workout) error
	DeleteWorkout(ctx context.Context, id uint) error
	ListExercises(ctx context.Context) ([]entity.Exercise, error)
	GetExerciseById(ctx context.Context, id uint) (*entity.Exercise, error)
	CreateDiet(ctx context.Context, diet *entity.Diet) error
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

func (r *repository) AddPhoto(ctx context.Context, photo *entity.Image) error {
	tx := r.db.WithContext(ctx).Create(photo)
	return tx.Error
}

func (r *repository) GetWorkoutById(ctx context.Context, id uint) (*entity.Workout, error) {
	workout := new(entity.Workout)
	tx := r.db.WithContext(ctx).Where(&entity.Workout{Model: gorm.Model{ID: id}}).First(workout)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return workout, tx.Error
}

func (r *repository) UpdateWorkout(ctx context.Context, workout entity.Workout) error {
	tx := r.db.WithContext(ctx).Save(workout)
	return tx.Error
}

func (r *repository) DeleteWorkout(ctx context.Context, id uint) error {
	tx := r.db.WithContext(ctx).Delete(&entity.Workout{}, id)
	return tx.Error
}

func (r *repository) ListExercises(ctx context.Context) ([]entity.Exercise, error) {
	var exercises []entity.Exercise
	tx := r.db.WithContext(ctx).Find(&exercises)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return exercises, nil
}

func (r *repository) GetExerciseById(ctx context.Context, id uint) (*entity.Exercise, error) {
	exercise := new(entity.Exercise)
	tx := r.db.WithContext(ctx).Find(&entity.Exercise{Model: gorm.Model{ID: id}}).First(exercise)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return exercise, tx.Error
}

func (r *repository) CreateDiet(ctx context.Context, diet *entity.Diet) error {
	tx := r.db.WithContext(ctx).Create(diet)
	return tx.Error
}
