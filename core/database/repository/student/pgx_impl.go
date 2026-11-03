package student_repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pact/core/database/repository"
)

var _ Repository = &pgxStudentRepository{}

type pgxStudentRepository struct {
	pool *pgxpool.Pool
}

func pgxPoolConnect(databaseURL string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func NewPgxStudentRepository(databaseURL string) (Repository, error) {
	pool, err := pgxPoolConnect(databaseURL)
	if err != nil {
		return nil, err
	}
	return &pgxStudentRepository{
		pool: pool,
	}, nil
}

func (sr *pgxStudentRepository) Create(ctx context.Context, params *CreateStudentParams) (*Student, error) {
	query := `
	 INSERT INTO users (name, email)
	 VALUES ($1, $2)
	 RETURNING *
	`
	row := sr.pool.QueryRow(ctx, query, params.Name, params.Email)
	var student Student
	err := row.Scan(
		&student.ID,
		&student.Name,
		&student.Email,
	)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (sr *pgxStudentRepository) GetByID(ctx context.Context, id string) (*Student, error) {
	query := `
		SELECT * 
		FROM users 
		WHERE id = $1
	`
	row := sr.pool.QueryRow(ctx, query, id)
	var student Student
	err := row.Scan(
		&student.ID,
		&student.Name,
		&student.Email,
	)
	if err == pgx.ErrNoRows {
		return nil, ErrStudentNotFound
	}
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (sr *pgxStudentRepository) GetUsers(ctx context.Context, params *GetUsersParams, pagination *repository.Pagination) (repository.PaginationResponse[Student], error) {
	result := repository.PaginationResponse[Student]{}
	query := `
		SELECT *
		FROM users
		WHERE 1=1
	`
	args := []any{}

	if params.Name != "" {
		args = append(args, "%"+params.Name+"%")
		query += fmt.Sprintf(" AND name ILIKE $%d", len(args))
	}

	args = append(args, pagination.Limit+1)
	query += fmt.Sprintf(" ORDER BY id LIMIT $%d", len(args))

	args = append(args, pagination.Page*pagination.Limit)
	query += fmt.Sprintf(" OFFSET $%d", len(args))

	rows, err := sr.pool.Query(ctx, query, args...)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	countQuery := `
		SELECT COUNT(*)
		FROM users
		WHERE 1=1
	`
	countArgs := []any{}
	if params.Name != "" {
		countArgs = append(countArgs, "%"+params.Name+"%")
		countQuery += fmt.Sprintf(" AND name ILIKE $%d", len(countArgs))
	}
	var total int

	err = sr.pool.QueryRow(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		return result, fmt.Errorf("count users: %w", err)
	}

	users := []Student{}
	for rows.Next() {
		var user Student

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		)
		if err != nil {
			return result, err
		}

		users = append(users, user)
	}

	hasNextPage := len(users) > pagination.Limit
	totalPages := (total + pagination.Limit - 1) / pagination.Limit
	result = repository.PaginationResponse[Student]{
		Data:        users,
		HasNextPage: hasNextPage,
		TotalPages:  totalPages,
	}

	return result, nil
}

func (sr *pgxStudentRepository) Close() {
	if sr.pool != nil {
		sr.pool.Close()
	}
}
