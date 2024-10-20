package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// User структура пользователя
type User struct {
	ID         int64  `json:"id"`
	KeycloakID string `json:"keycloak_id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Version    int    `json:"version"`
}

// UserRepository интерфейс для работы с таблицей users
type UserRepository interface {
	CreateUser(ctx context.Context, u *User) error
	GetUserByKeycloakID(ctx context.Context, keycloakID string) (*User, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	UpdateUser(ctx context.Context, u *User) error
	DeleteUser(ctx context.Context, id int64) error
}

type userRepository struct {
	db *pgxpool.Pool
}

// NewUserRepository создает новый экземпляр userRepository
func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

// CreateUser создает нового пользователя в базе
func (r *userRepository) CreateUser(ctx context.Context, user *User) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO users (keycloak_id, name, surname, email, role, created_at, updated_at, version)
         VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		user.KeycloakID, user.Name, user.Surname, user.Email, user.Role, time.Now(), time.Now(), 1,
	)
	return err
}

// GetUserByKeycloakID находит пользователя по keycloak_id
func (r *userRepository) GetUserByKeycloakID(ctx context.Context, keycloakID string) (*User, error) {
	var user User
	err := r.db.QueryRow(
		ctx,
		`SELECT id, keycloak_id, name, surname, email, role, created_at, updated_at, version FROM users WHERE keycloak_id=$1`,
		keycloakID,
	).Scan(&user.ID, &user.KeycloakID, &user.Name, &user.Surname, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

// GetUserByID находит пользователя по id
func (r *userRepository) GetUserByID(ctx context.Context, id int64) (*User, error) {
	var user User
	err := r.db.QueryRow(
		ctx,
		`SELECT id, keycloak_id, name, surname, email, role, created_at, updated_at, version FROM users WHERE keycloak_id=$1`,
		id,
	).Scan(&user.ID, &user.KeycloakID, &user.Name, &user.Surname, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

// GetUserByEmail находит пользователя по email
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := r.db.QueryRow(
		ctx,
		`SELECT id, keycloak_id, name, surname, email, role, created_at, updated_at, version FROM users WHERE email=$1`,
		email,
	).Scan(&user.ID, &user.KeycloakID, &user.Name, &user.Surname, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

// GetAllUsers возвращает список всех пользователей
func (r *userRepository) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := r.db.Query(ctx, `SELECT id, keycloak_id, name, surname, email, role, created_at, updated_at, version FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.KeycloakID, &user.Name, &user.Surname, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// UpdateUser обновляет информацию о пользователе
func (r *userRepository) UpdateUser(ctx context.Context, user *User) error {
	_, err := r.db.Exec(
		ctx,
		`UPDATE users SET name=$1, surname=$2, email=$3, role=$4, updated_at=$5 WHERE keycloak_id=$6`,
		user.Name, user.Surname, user.Email, user.Role, time.Now(), user.KeycloakID,
	)
	return err
}

// DeleteUser удаляет пользователя по keycloak_id
func (r *userRepository) DeleteUser(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, `DELETE FROM users WHERE keycloak_id=$1`, id)
	return err
}
