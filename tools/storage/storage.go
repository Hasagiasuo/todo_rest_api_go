package storage

import (
	"errors"
	"fmt"
	"log/slog"
	"pet_pr/tools/configs"
	"pet_pr/tools/models"
	"strconv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	DataBase *sqlx.DB
	Logger *slog.Logger
}

func InitStorage(cfg configs.DBConfig, logger *slog.Logger) *Storage {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Name, cfg.DB_name, cfg.SSL_mode))
	if err != nil {
		logger.Error(fmt.Sprintf("Database not connected: %s", err.Error()))
		return nil
	}
	err = db.Ping();
	if err != nil {
		logger.Error(fmt.Sprintf("Database not responce: %s", err.Error()))
		return nil
	}
	logger.Info("DataBase connected")
	return &Storage{ DataBase: db, Logger: logger }
}

func (storage Storage) CreateTables() {
	storage.CreateTodoListTables()
	storage.CreateUserTables()
}

func (storage Storage) CreateTodoListTables() {
	query := `
		CREATE TABLE IF NOT EXISTS todo (
			id SERIAL PRIMARY KEY,
			uid INT,
			title VARCHAR(50),
			done BOOLEAN
		);
	`
	_, err := storage.DataBase.Exec(query)
	if err != nil {
		storage.Logger.Info(fmt.Sprintf("Error exec: %s", err.Error()))
		return
	}
	storage.Logger.Info("Table todo created")
}

func (storage Storage) CreateUserTables() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(20),
			email VARCHAR(30),
			password VARCHAR(20)
		);
	`
	_, err := storage.DataBase.Exec(query)
	if err != nil {
		storage.Logger.Info(fmt.Sprintf("Error exec: %s", err.Error()))
		return
	}
	storage.Logger.Info("Table user created")
}

func (storage Storage) GetUserById(id string) models.User {
	user := models.User{}
	query := `SELECT id, name, email, password FROM users WHERE id=$1`
	if err := storage.DataBase.Get(&user, query, id); err != nil {
		storage.Logger.Info("Cannot get user from database")
	}
	return user
}

func (storage Storage) GetUserByEmail(email string) models.User {
	user := models.User{}
	query := `SELECT id, name, email, password FROM users WHERE email=$1`
	if err := storage.DataBase.Get(&user, query, email); err != nil {
		storage.Logger.Info("Cannot get user from database")
	}
	return user
}

func (storage Storage) GetUserByName(name string) models.User {
	user := models.User{}
	query := `SELECT id, name, email, password FROM users WHERE name=$1`
	if err := storage.DataBase.Get(&user, query, name); err != nil {
		storage.Logger.Info("Cannot get user from database")
	}
	return user
}

func (storage Storage) GetUserTodosByUID(uid string) []models.TodoItem {
	var result []models.TodoItem
	query := `SELECT * FROM todo WHERE uid=$1`
	if err := storage.DataBase.Select(&result, query, uid); err != nil {
		storage.Logger.Info("Cannot get todo by uid")
	}
	return result
}

func (storage Storage) UpdateDoneTask(n_val bool, id string) {
	query := `UPDATE todo SET done = $1 WHERE id = $2;`
	_, err := storage.DataBase.Exec(query, n_val, id);
	if err != nil {
		storage.Logger.Info("Cannot update todo item")
		return
	}
}

func (storage Storage) GetIDByTitle(title string) string {
	var id string
	query := `SELECT id FROM todo WHERE title=$1`
	if err := storage.DataBase.Get(&id, query, title); err != nil {
		storage.Logger.Info("Cannot get todo item")
		return ""
	}
	return id
}

func (storage Storage) PushTodoByUID(title string, uid string) error {
	var tmp string = ""
	get_q := `SELECT title FROM todo WHERE title=$1`
	storage.DataBase.Get(&tmp, get_q, title)
	if tmp != "" {
		storage.Logger.Info("Item already exists")
		return errors.New("ITEM ALREADY EXISTS")
	}
	query := `INSERT INTO todo (uid, title, done) VALUES ($1, $2, false);`
	i_uid, err := strconv.Atoi(uid)
	if err != nil {
		storage.Logger.Info("Cannot convert UID")
		return errors.New("CANNOT CONVERT UID")
	}
	if _, err := storage.DataBase.Exec(query, i_uid, title); err != nil {
		storage.Logger.Info(fmt.Sprintf("Cannot insert new todo item: %s", err.Error()))
		return errors.New("CANNOT INSERT NEW TODO")
	}
	return nil
}

func (storage Storage) DeleteTodoByID(id string) error {
	query := `DELETE FROM todo WHERE id=$1;`
	if _, err := storage.DataBase.Exec(query, id); err != nil {
		storage.Logger.Info(fmt.Sprintf("Cannot delete item: %s", err.Error()))
		return errors.New("CANNOT DELETE")
	}
	return nil
}

func (storage Storage) CreateNewUser(name string, email string, password string) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`
	if _, err := storage.DataBase.Exec(query, name, email, password); err != nil {
		storage.Logger.Info(fmt.Sprintf("Cannot push user: %s", err.Error()))
		return errors.New("CANNOT PUSH USER")
	}
	return nil
}