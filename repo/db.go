package repo

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/AllenKaplan/hackpack-go/models"
)

type Database struct {
	db *gorm.DB
}

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func NewDB(dbCfg Config) (Repository, error) {
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbCfg.Host, dbCfg.User, dbCfg.Password, dbCfg.DBName, dbCfg.Port)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(models.Todo{})
	return Database{db: db}, nil
}

type Repository interface {
	GetTodo(context.Context, uint) (models.Todo, error)
	GetTodos(context.Context) ([]models.Todo, error)
	CreateTodo(context.Context, models.Todo) (models.Todo, error)
	UpdateTodo(context.Context, models.Todo) (models.Todo, error)
	DeleteTodo(context.Context, uint) (bool, error)
}

func (db Database) GetTodo(ctx context.Context, id uint) (models.Todo, error) {
	var app models.Todo
	result := db.db.First(&app, "id = ?", id)
	if result.Error != nil {
		return models.Todo{}, result.Error
	}
	return app, nil
}

func (db Database) GetTodos(ctx context.Context) ([]models.Todo, error) {
	var apps []models.Todo
	result := db.db.Find(&apps)
	if result.Error != nil {
		return nil, result.Error
	}
	return apps, nil
}

func (db Database) CreateTodo(ctx context.Context, app models.Todo) (models.Todo, error) {
	result := db.db.Clauses(clause.Returning{}).Create(app)
	if result.Error != nil {
		return models.Todo{}, result.Error
	}
	return app, nil
}

func (db Database) UpdateTodo(ctx context.Context, app models.Todo) (models.Todo, error) {
	result := db.db.Clauses(clause.Returning{}).Updates(app)
	if result.Error != nil {
		return models.Todo{}, nil
	}
	return app, nil

}

func (db Database) DeleteTodo(ctx context.Context, id uint) (bool, error) {
	result := db.db.Delete(&models.Todo{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
