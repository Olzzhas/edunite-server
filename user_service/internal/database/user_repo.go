package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Surname   string
	Email     string
	Role      string
	CreatedAt string
	UpdatedAt string
	Version   int
}

type UserRepository interface {
	CreateUser(ctx context.Context, u *User) error
	GetUserByID(ctx context.Context, id int64) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	UpdateUser(ctx context.Context, u *User) error
	DeleteUser(ctx context.Context, id int64) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

// CreateUser создает нового пользователя с переданным ID
func (r *userRepository) CreateUser(ctx context.Context, user *User) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO users (id, name, surname, email, role, created_at, updated_at, version)
         VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		user.ID, user.Name, user.Surname, user.Email, user.Role, time.Now(), time.Now(), 1,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByID возвращает пользователя по ID
func (r *userRepository) GetUserByID(ctx context.Context, id int64) (*User, error) {
	var user User
	err := r.db.QueryRow(
		ctx,
		`SELECT id, name, surname, email, role, created_at, updated_at, version FROM users WHERE id=$1`,
		id,
	).Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.Version)
	return &user, err
}

// GetUserByEmail возвращает пользователя по email
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := r.db.QueryRow(
		ctx,
		`SELECT id, name, surname, email, role, created_at, updated_at, version FROM users WHERE email=$1`,
		email,
	).Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.Version)
	return &user, err
}

// GetAllUsers возвращает всех пользователей
func (r *userRepository) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, surname, email, role, created_at, updated_at, version FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.Version); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// UpdateUser обновляет данные пользователя
func (r *userRepository) UpdateUser(ctx context.Context, user *User) error {
	_, err := r.db.Exec(
		ctx,
		`UPDATE users SET name=$1, surname=$2, email=$3, role=$4, updated_at=$5 WHERE id=$6`,
		user.Name, user.Surname, user.Email, user.Role, time.Now(), user.ID,
	)
	return err
}

// DeleteUser удаляет пользователя по ID
func (r *userRepository) DeleteUser(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, `DELETE FROM users WHERE id=$1`, id)
	return err
}
