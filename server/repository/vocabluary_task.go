package repository

import (
	"DB_course_paper/server/entity"
	"context"
	"database/sql"
)

func (r *repository) CreateOptionsTask(ctx context.Context, task entity.VocabularyOptionsTask) (int, error) {
	query := `
        INSERT INTO vocabulary_options_task (sentence, answer_options, explanation)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	var id int
	err := r.db.QueryRowContext(ctx, query,
		task.Sentence,
		task.AnswerOptions,
		task.Explanation,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) DeleteOptionsTask(ctx context.Context, taskId int) error {
	query := `
        DELETE FROM vocabulary_options_task 
        WHERE id = $1
    `

	result, err := r.db.ExecContext(ctx, query, taskId)
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

func (r *repository) GetOptionsTaskById(ctx context.Context, taskId int) (entity.VocabularyOptionsTask, error) {
	query := `
        SELECT id, sentence, answer_options, explanation 
        FROM vocabulary_options_task 
        WHERE id = $1
    `

	var task entity.VocabularyOptionsTask
	err := r.db.QueryRowContext(ctx, query, taskId).Scan(
		&task.ID,
		&task.Sentence,
		&task.AnswerOptions,
		&task.Explanation,
	)

	if err == sql.ErrNoRows {
		return task, nil
	} else if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) CreateWordTask(ctx context.Context, task entity.VocabularyWordTask) (int, error) {
	query := `
        INSERT INTO vocabulary_word_task (sentence, answer, explanation)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	var id int
	err := r.db.QueryRowContext(ctx, query,
		task.Sentence,
		task.Answer,
		task.Explanation,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) DeleteWordTask(ctx context.Context, taskId int) error {
	query := `
        DELETE FROM vocabulary_word_task 
        WHERE id = $1
    `

	result, err := r.db.ExecContext(ctx, query, taskId)
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

func (r *repository) GetWordTaskById(ctx context.Context, taskId int) (entity.VocabularyWordTask, error) {
	query := `
        SELECT id, sentence, answer, explanation 
        FROM vocabulary_word_task 
        WHERE id = $1
    `

	var task entity.VocabularyWordTask
	err := r.db.QueryRowContext(ctx, query, taskId).Scan(
		&task.ID,
		&task.Sentence,
		&task.Answer,
		&task.Explanation,
	)

	if err == sql.ErrNoRows {
		return task, nil
	} else if err != nil {
		return task, err
	}

	return task, nil
}
