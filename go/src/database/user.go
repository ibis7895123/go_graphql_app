package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (user *User) TableName() string {
	return "user"
}

type UserDaoInterface interface {
	InsertOne(user *User) error
	FindAll() ([]*User, error)
	FindOne(id string) (*User, error)
	FindByTodoID(todoID string) (*User, error)
	ExistUserID(id string) error
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
	response := userDao.db.
		Order("updated_at desc").
		Find(&users)

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
	var user User
	response := userDao.db.
		Where("id = ?", id).
		First(&user)

	// DBエラー
	// データ無しの場合もここに入る
	if err := response.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

/**
 * todoIDからユーザーを検索
 */
func (userDao *UserDao) FindByTodoID(todoID string) (*User, error) {
	var user User

	response := userDao.db.Table("user").
		Select("user.*").
		Joins("LEFT JOIN todo ON todo.user_id = user.id").
		Where("todo.id = ?", todoID).
		First(&user)

	// DBエラー
	// データ無しの場合もここに入る
	if err := response.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

/**
 * ユーザーIDが存在するかチェック
 */
func (userDao *UserDao) ExistUserID(id string) error {
	var user User

	response := userDao.db.
		Where("id = ?", id).
		First(&user)

	// DBエラー
	// データ無しの場合もここに入る
	if err := response.Error; err != nil {
		return err
	}

	return nil
}
