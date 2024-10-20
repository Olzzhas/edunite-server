package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type User struct {
	ID        int
	Name      string
	Surname   string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

type UserRepository interface {
	CreateUser(ctx context.Context, u *User) (int, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	UpdateUser(ctx context.Context, u *User) error
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

// CreateUser Создать пользователя
func (r *userRepository) CreateUser(ctx context.Context, user *User) (int, error) {
	var id int
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO users (name, surname, role, created_at, updated_at, version)
         VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		user.Name, user.Surname, user.Role, time.Now(), time.Now(), 1,
	).Scan(&id)
	return id, err
}

// GetUserByID Получить пользователя по ID
func (r *userRepository) GetUserByID(ctx context.Context, id int) (*User, error) {
	var user User
	err := r.db.QueryRow(
		ctx,
		`SELECT id, name, surname, role, created_at, updated_at, version FROM users WHERE id=$1`,
		id,
	).Scan(&user.ID, &user.Name, &user.Surname, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.Version)
	return &user, err
}

// GetUserByEmail Получить пользователя по Email
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := r.db.QueryRow(
		ctx,
		`SELECT id, name, surname, role, created_at, updated_at, version FROM users WHERE email=$1`,
		email,
	).Scan(&user.ID, &user.Name, &user.Surname, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.Version)
	return &user, err
}

// GetAllUsers Получить всех пользователей
func (r *userRepository) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, surname, role, created_at, updated_at, version FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.Version); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// UpdateUser Обновить пользователя
func (r *userRepository) UpdateUser(ctx context.Context, user *User) error {
	_, err := r.db.Exec(
		ctx,
		`UPDATE users SET name=$1, surname=$2, role=$3, updated_at=$4 WHERE id=$5`,
		user.Name, user.Surname, user.Role, time.Now(), user.ID,
	)
	return err
}

// DeleteUser Удалить пользователя
func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM users WHERE id=$1`, id)
	return err
}
