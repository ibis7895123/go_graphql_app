package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID   string `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`
}

func (user *User) TableName() string {
	return "user"
}

type UserDaoInterface interface {
	InsertOne(user *User) error
	FindAll() ([]*User, error)
	FindOne(id string) (*User, error)
	FindByTodoID(todoID string) (*User, error)
}

type UserDao struct {
	db *gorm.DB
}

/**
 * database access objectの生成
 */
func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

/**
 * 1レコードを挿入
 */
func (userDao *UserDao) InsertOne(user *User) error {
	response := userDao.db.Create(user)

	// DBエラー
	if err := response.Error; err != nil {
		return err
	}

	return nil
}

/**
 * 全ユーザーを取得
 */
func (userDao *UserDao) FindAll() ([]*User, error) {
	var users []*User
	// ユーザー全取得
	response := userDao.db.Find(&users)

	// DBエラー
	if err := response.Error; err != nil {
		return nil, err
	}

	return users, nil
}

/**
 * 1ユーザーを検索
 */
func (userDao *UserDao) FindOne(id string) (*User, error) {
	var users []*User
	response := userDao.db.Where("id = ?", id).Find(&users)

	// DBエラー
	if err := response.Error; err != nil {
		return nil, err
	}

	// データなし
	if len(users) == 0 {
		return nil, fmt.Errorf("not found user")
	}

	return users[0], nil
}

func (userDao *UserDao) FindByTodoID(todoID string) (*User, error) {
	var users []*User

	response := userDao.db.Table("user").
		Select("user.*").
		Joins("LEFT JOIN todo ON todo.user_id = user.id").
		Where("todo.id = ?", todoID).
		First(&users)

	// DBエラー
	if err := response.Error; err != nil {
		return nil, err
	}

	// データなし
	if len(users) == 0 {
		return nil, fmt.Errorf("not found user")
	}

	return users[0], nil
}
