package database

import (
	"context"
	"errors"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"sort"
	"strconv"
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
	AddPhoto(ctx context.Context, photo *entity.Photo) error
	GetWorkoutById(ctx context.Context, id uint) (*entity.Workout, error)
	UpdateWorkout(ctx context.Context, workout entity.Workout) error
	DeleteWorkout(ctx context.Context, id uint) error
	ListExercises(ctx context.Context) ([]entity.Exercise, error)
	GetExerciseById(ctx context.Context, id uint) (*entity.Exercise, error)
	GetImages(ctx context.Context) ([]entity.Photo, error)
	GetStudentsOfCoach(ctx context.Context, coachId uint) ([]entity.User, error)
	CreateDiet(ctx context.Context, diet *entity.Diet) error
	UpdateDiet(ctx context.Context, diet *entity.Diet) error
	GetDietByID(ctx context.Context, dietID uint) (*entity.Diet, error)
	DeleteDiet(ctx context.Context, dietId uint) error
	ListDiets(ctx context.Context, userId uint) ([]entity.Diet, error)
	ListFoods(ctx context.Context) ([]*entity.Food, error)
	DeletePhotoById(ctx context.Context, id uint) error
	DeletePrevProfilePictures(ctx context.Context, userId uint) error
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

func (r *repository) AddPhoto(ctx context.Context, photo *entity.Photo) error {
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

func (r *repository) GetImages(ctx context.Context) ([]entity.Photo, error) {
	var images []entity.Photo
	tx := r.db.WithContext(ctx).Find(&images)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return images, nil
}

func (r *repository) GetStudentsOfCoach(ctx context.Context, coachId uint) ([]entity.User, error) {
	workouts, err := r.ListWorkoutByCoachId(ctx, coachId)
	if err != nil {
		return nil, err
	}

	userIds := lo.Uniq(lo.Map(workouts, func(w entity.Workout, _ int) uint {
		return w.UserId
	}))

	users, err := r.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	usersFiltered := lo.Filter(users, func(u entity.User, i int) bool {
		return lo.Contains(userIds, u.ID)
	})

	return usersFiltered, nil
}

func (r *repository) CreateDiet(ctx context.Context, diet *entity.Diet) error {
	return r.db.WithContext(ctx).Create(diet).Error
}

func (r *repository) UpdateDiet(ctx context.Context, diet *entity.Diet) error {
	return r.db.WithContext(ctx).Save(diet).Error
}

func (r *repository) GetDietByID(ctx context.Context, dietID uint) (*entity.Diet, error) {
	var diet entity.Diet
	if err := r.db.WithContext(ctx).Preload("Meals").First(&diet, dietID).Error; err != nil {
		return nil, err
	}
	return &diet, nil
}
func (r *repository) DeleteDiet(ctx context.Context, dietId uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Diet{}, dietId).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) ListDiets(ctx context.Context, userId uint) ([]entity.Diet, error) {
	var diets []entity.Diet
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Preload("Meals").Find(&diets).Error; err != nil {
		return nil, err
	}
	return diets, nil
}

func (r *repository) ListFoods(ctx context.Context) ([]*entity.Food, error) {
	var foods []*entity.Food
	if err := r.db.WithContext(ctx).Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}

func (r *repository) DeletePhotoById(ctx context.Context, id uint) error {
	tx := r.db.WithContext(ctx).Delete(&entity.Photo{}, id)
	return tx.Error
}

func (r *repository) DeletePrevProfilePictures(ctx context.Context, userId uint) error {
	prefix := "pp_" + strconv.Itoa(int(userId)) + "%"
	var profilePhotos []*entity.Photo
	tx := r.db.WithContext(ctx).Find(&profilePhotos, "photo_name LIKE ?", prefix)
	if tx.Error != nil {
		return tx.Error
	}

	sort.Slice(profilePhotos, func(i, j int) bool {
		return profilePhotos[i].CreatedAt.After(profilePhotos[j].CreatedAt)
	})

	if len(profilePhotos) > 1 {
		tx = r.db.WithContext(ctx).Delete(profilePhotos[1:])
	}

	return tx.Error
}
