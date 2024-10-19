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
	UpdateUser(ctx context.Context, u *User) (*User, error)
	DeleteUser(ctx context.Context, id int) (int, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

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

func (r *userRepository) GetUserByID(ctx context.Context, id int) (u *User, err error) {
	err = r.db.QueryRow(
		ctx,
		`SELECT id, name, surname, role, created_at, updated_at, version FROM users WHERE id = $1`,
		id,
	).Scan(&u.ID, &u.Name, &u.Surname, &u.Role, &u.CreatedAt, &u.UpdatedAt, &u.Version)

	return u, err
}

func (r *userRepository) UpdateUser(ctx context.Context, user *User) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) (int, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}
