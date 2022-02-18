package repo

import (
	"context"
	"errors"

	"github.com/AllenKaplan/hackpack-go/models"
)

type FakeDatabase struct {
	db map[uint]models.Todo
}

func newFakeDB() Repository {
	return FakeDatabase{db: make(map[uint]models.Todo)}
}

func (db FakeDatabase) GetTodo(ctx context.Context, id uint) (models.Todo, error) {
	app, ok := db.db[id]
	if !ok {
		return models.Todo{}, errors.New("Todo does not exist")
	}
	return app, nil
}

func (db FakeDatabase) GetTodos(ctx context.Context) ([]models.Todo, error) {
	var apps []models.Todo
	for _, app := range db.db {
		apps = append(apps, app)
	}
	return apps, nil
}

func (db FakeDatabase) CreateTodo(ctx context.Context, app models.Todo) (models.Todo, error) {
	db.db[uint(len(db.db)+1)] = app
	return app, nil
}

func (db FakeDatabase) UpdateTodo(ctx context.Context, app models.Todo) (models.Todo, error) {
	db.db[app.ID] = app
	return app, nil
}

func (db FakeDatabase) DeleteTodo(ctx context.Context, id uint) (bool, error) {
	delete(db.db, id)
	return true, nil
}
