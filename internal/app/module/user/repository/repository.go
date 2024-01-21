package repository

import (
	"CodeWithAzri/internal/app/module/user/entity"
	"database/sql"
)

type UserRepository interface {
	Create(e *entity.User) error
	ReadMany(limit, offset int) ([]entity.User, error)
	ReadOne(id string) (*entity.User, error)
	Update(id string, e *entity.User) error
	Delete(id string) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	r := &Repository{db: db}
	return r
}

func (r *Repository) Create(e *entity.User) error {
	query := "INSERT INTO users (id, name, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, e.ID, e.Name, e.Email, e.CreatedAt, e.UpdatedAt)
	return err
}

func (r *Repository) ReadMany(limit, offset int) ([]entity.User, error) {
	query := "SELECT * FROM users LIMIT $1 OFFSET $2"
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *Repository) ReadOne(id string) (*entity.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) Update(id string, e *entity.User) error {
	query := "UPDATE users SET name = $1 WHERE id = $2"
	_, err := r.db.Exec(query, e.Name, id)
	return err
}

func (r *Repository) Delete(id string) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
