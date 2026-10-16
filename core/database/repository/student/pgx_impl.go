package student_repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var _ StudentRepository = &pgxStudentRepository{}

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

func NewPgxStudentRepository(databaseURL string) (StudentRepository, error) {
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
	 INSERT INTO students (name, email)
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
		FROM students 
		WHERE id = $1
	`
	row := sr.pool.QueryRow(ctx, query, id)
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

func (sr *pgxStudentRepository) Close() {
	if sr.pool != nil {
		sr.pool.Close()
	}
}
