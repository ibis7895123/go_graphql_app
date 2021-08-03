package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	ID        string    `gorm:"column:id;primary_key"`
	Text      string    `gorm:"column:text"`
	Done      bool      `gorm:"column:done"`
	UserID    string    `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (todo *Todo) TableName() string {
	return "todo"
}

type TodoDaoInterface interface {
	InsertOne(todo *Todo) error
	FindAll() ([]*Todo, error)
	FindOne(id string) (*Todo, error)
	FindByUserID(userID string) ([]*Todo, error)
}

type TodoDao struct {
	db *gorm.DB
}

/**
 * database access objectの生成
 */
func NewTodoDao(db *gorm.DB) *TodoDao {
	return &TodoDao{db: db}
}

/**
 * 1レコードを挿入
 */
func (todoDao *TodoDao) InsertOne(todo *Todo) error {
	response := todoDao.db.Create(todo)

	// DBエラー
	if err := response.Error; err != nil {
		return err
	}

	return nil
}

/**
 * 全todoを取得
 */
func (todoDao *TodoDao) FindAll() ([]*Todo, error) {
	var todos []*Todo

	response := todoDao.db.
		Order("updated_at desc").
		Find(&todos)

	// DBエラー
	// データ無しの場合もここに入る
	if err := response.Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (todoDao *TodoDao) FindOne(id string) (*Todo, error) {
	var todo Todo

	response := todoDao.db.
		Where("id = ?", id).
		First(&todo)

	// DBエラー
	// データ無しの場合もここに入る
	if err := response.Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (todoDao *TodoDao) FindByUserID(userID string) ([]*Todo, error) {
	var todos []*Todo

	response := todoDao.db.
		Where("user_id = ?", userID).
		Find(&todos)

	// DBエラー
	// データ無しの場合もここに入る
	if err := response.Error; err != nil {
		return nil, err
	}

	return todos, nil
}
