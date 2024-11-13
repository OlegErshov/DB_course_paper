package repository

import (
	"DB_course_paper/server/entity"
	"context"
	"database/sql"
	"time"
)

func (r repository) CreateTeacher(ctx context.Context, teacher entity.Teacher) (int, error) {
	query := `
        INSERT INTO teachers (name, email, phone, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `

	var id int
	err := r.db.QueryRowContext(ctx, query,
		teacher.Name,
		teacher.Email,
		teacher.Phone,
		teacher.Password,
		time.Now(), // created_at
		time.Now(), // updated_at
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r repository) DeleteTeacher(ctx context.Context, teacherId int) error {
	query := `
        DELETE FROM teachers 
        WHERE id = $1
    `

	result, err := r.db.ExecContext(ctx, query, teacherId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows // Ошибка, если не найдено ни одной записи
	}

	return nil
}

func (r repository) UpdateTeacher(ctx context.Context, teacher entity.Teacher) (int, error) {
	query := `
        UPDATE teachers 
        SET name = $1, email = $2, phone = $3, password = $4, updated_at = $5
        WHERE id = $6
        RETURNING id
    `

	var updatedID int
	err := r.db.QueryRowContext(ctx, query,
		teacher.Name,
		teacher.Email,
		teacher.Phone,
		teacher.Password,
		time.Now(), // обновляем поле updated_at на текущее время
		teacher.ID, // ID учителя, которого обновляем
	).Scan(&updatedID)

	if err != nil {
		return 0, err
	}

	return updatedID, nil
}

func (r repository) GetStudents(ctx context.Context, teacherId int) ([]entity.Student, error) {
	query := `
        SELECT s.id, s.name, s.email, s.phone, s.password, s.created_at, s.updated_at 
        FROM students s
        INNER JOIN groups g ON s.id = g.teacher_id
        WHERE g.teacher_id = $1
    `

	rows, err := r.db.QueryContext(ctx, query, teacherId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []entity.Student
	for rows.Next() {
		var student entity.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.Phone, &student.Password, &student.CreatedAt, &student.UpdatedAt)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r repository) GetTeacherTopics(ctx context.Context, teacherId int) ([]entity.Topic, error) {
	query := `
        SELECT t.id, t.name, t.mark 
        FROM topic t
        INNER JOIN groups g ON t.id = g.teacher_id
        WHERE g.teacher_id = $1
    `

	rows, err := r.db.QueryContext(ctx, query, teacherId)
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

func (r repository) GetTeacherByCreds(ctx context.Context, phone, password string) (entity.Teacher, error) {
	var teacher entity.Teacher
	query := `SELECT id, name, email, phone, password, created_at, updated_at 
              FROM teachers 
              WHERE phone = $1 AND password = $2`

	err := r.db.QueryRowContext(ctx, query, phone, password).Scan(
		&teacher.ID,
		&teacher.Name,
		&teacher.Email,
		&teacher.Phone,
		&teacher.Password,
		&teacher.CreatedAt,
		&teacher.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Teacher{}, nil // Teacher not found
		}
		return entity.Teacher{}, err // Database error
	}

	return teacher, nil
}
