package repository

import (
	"DB_course_paper/server/entity"
	"context"
	"database/sql"
)

func (r *repository) CreateTopic(ctx context.Context, topic entity.Topic) (int, error) {
	query := `
        INSERT INTO topic (name, mark)
        VALUES ($1, $2)
        RETURNING id
    `

	var id int
	err := r.db.QueryRowContext(ctx, query,
		topic.Name,
		topic.Mark,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) DeleteTopic(ctx context.Context, topicId int) error {
	query := `
        DELETE FROM topic 
        WHERE id = $1
    `

	result, err := r.db.ExecContext(ctx, query, topicId)
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

func (r *repository) GetTopicById(ctx context.Context, topicId int) (entity.Topic, error) {
	query := `
        SELECT id, name, mark 
        FROM topic 
        WHERE id = $1
    `

	var topic entity.Topic
	err := r.db.QueryRowContext(ctx, query, topicId).Scan(
		&topic.ID,
		&topic.Name,
		&topic.Mark,
	)

	if err == sql.ErrNoRows {
		return topic, nil // Возвращаем пустую запись, если тема не найдена
	} else if err != nil {
		return topic, err
	}

	return topic, nil
}

func (r *repository) UpdateTopic(ctx context.Context, topic entity.Topic) (int, error) {
	query := `
        UPDATE topic
        SET name = $1, mark = $2
        WHERE id = $3
        RETURNING id
    `

	var id int
	err := r.db.QueryRowContext(ctx, query,
		topic.Name,
		topic.Mark,
		topic.ID,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
