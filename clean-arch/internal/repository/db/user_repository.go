package db

import (
	"clean-arch/internal/domain"
	"clean-arch/internal/model/entity"
	"context"
	"database/sql"
	"errors"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {

	query := `
		INSERT INTO users (name, email)
		VALUES (?, ?)
	`

	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)

	return nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {

	query := `
		SELECT id, name, email
		FROM users
		ORDER BY id
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User

	for rows.Next() {

		var user entity.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*entity.User, error) {

	query := `
		SELECT id, name, email
		FROM users
		WHERE id = ?
	`

	var user entity.User

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {

	query := `
		UPDATE users
		SET
			name = ?,
			email = ?
		WHERE id = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.ID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {

	query := `
		DELETE FROM users
		WHERE id = ?
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}