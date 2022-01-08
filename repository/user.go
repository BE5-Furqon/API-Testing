package repository

import (
	"layerarch/constant"
	"layerarch/entities"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

type User interface {
	GetAll() ([]entities.User, error)
	Get(userId int) (entities.User, error)
	Register(entities.User) (entities.User, error)
	Edit(user entities.User, userId int) (entities.User, error)
	Delete(userId int) (entities.User, error)
	Login(email, password string) (interface{}, error)
}

func (ur *UserRepository) GetAll() ([]entities.User, error) {
	var users []entities.User

	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) Get(userId int) (entities.User, error) {
	var user entities.User

	if err := ur.db.First(&user, userId).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Register(user entities.User) (entities.User, error) {
	if err := ur.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Edit(user entities.User, userId int) (entities.User, error) {
	var u entities.User
	ur.db.First(&u, userId)

	if err := ur.db.Model(&u).Updates(user).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (ur *UserRepository) Delete(userId int) (entities.User, error) {
	var user entities.User

	if err := ur.db.Delete(&user, userId).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Login(email, password string) (interface{}, error) {
	var user entities.User	

	if err := ur.db.First(&user, "email = ? and password = ?", email, password).Error; err != nil {
		return nil, err
	}

	ur.db.Model(&user).Updates(entities.User{
		Name:           user.Name,
		Email:          user.Email,
		Password:       user.Password,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 72).Unix()},
	})

	// create token with payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)

	// generate token
	t, err := token.SignedString([]byte(constant.JWT_SECRET))

	if err != nil {
		return nil, err
	}

	return t, nil
}