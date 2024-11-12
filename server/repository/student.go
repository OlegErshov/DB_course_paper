package repository

import (
	"DB_course_paper/server/entity"
	"context"
	"database/sql"
	"time"
)

func (r *repository) CreateStudent(ctx context.Context, student entity.Student) (int, error) {
	query := `
        INSERT INTO students (name, email, phone, password)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `

	var id int
	err := r.db.QueryRowContext(ctx, query,
		student.Name,
		student.Email,
		student.Phone,
		student.Password,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) DeleteStudent(ctx context.Context, studentId int) error {
	query := `
        DELETE FROM students 
        WHERE id = $1
    `

	result, err := r.db.ExecContext(ctx, query, studentId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *repository) UpdateStudent(ctx context.Context, student entity.Student) (int, error) {
	query := `
        UPDATE students 
        SET name = $1, email = $2, phone = $3, password = $4, updated_at = $5
        WHERE id = $6
        RETURNING id
    `

	var updatedID int
	err := r.db.QueryRowContext(ctx, query,
		student.Name,
		student.Email,
		student.Phone,
		student.Password,
		time.Now(),
		student.ID,
	).Scan(&updatedID)

	if err != nil {
		return 0, err
	}

	return updatedID, nil
}

func (r *repository) GetStudentById(ctx context.Context, studentId int) (entity.Student, error) {
	query := `
        SELECT id, name, email, phone, password, created_at, updated_at 
        FROM students 
        WHERE id = $1
    `

	var student entity.Student
	err := r.db.QueryRowContext(ctx, query, studentId).Scan(
		&student.ID,
		&student.Name,
		&student.Email,
		&student.Phone,
		&student.Password,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return student, nil
	} else if err != nil {
		return student, err
	}

	return student, nil
}

func (r *repository) GetStudentTopics(ctx context.Context, studentId int) ([]entity.Topic, error) {
	query := `
        SELECT t.id, t.name, t.mark 
        FROM topic t
        INNER JOIN topic_student ts ON t.id = ts.topic_id
        WHERE ts.student_id = $1
    `

	rows, err := r.db.QueryContext(ctx, query, studentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topics []entity.Topic
	for rows.Next() {
		var topic entity.Topic
		err := rows.Scan(&topic.ID, &topic.Name, &topic.Mark)
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return topics, nil
}
