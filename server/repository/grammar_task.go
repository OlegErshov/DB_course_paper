package repository

import (
	"DB_course_paper/server/entity"
	"context"
	"database/sql"
)

func (r *repository) CreateGrammarTask(ctx context.Context, task entity.GrammarTask) (int, error) {
	query := `
        INSERT INTO grammar_task (sentence, right_answer, hint, explanation)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `

	var id int
	err := r.db.QueryRowContext(ctx, query,
		task.Sentence,
		task.RightAnswer,
		task.Hint,
		task.Explanation,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) DeleteGrammarTask(ctx context.Context, taskId int) error {
	query := `
        DELETE FROM grammar_task 
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
		return sql.ErrNoRows // Ошибка, если задача с указанным ID не найдена
	}

	return nil
}

func (r *repository) GetGrammarTaskById(ctx context.Context, taskId int) (entity.GrammarTask, error) {
	query := `
        SELECT id, sentence, right_answer, hint, explanation 
        FROM grammar_task 
        WHERE id = $1
    `

	var task entity.GrammarTask
	err := r.db.QueryRowContext(ctx, query, taskId).Scan(
		&task.ID,
		&task.Sentence,
		&task.RightAnswer,
		&task.Hint,
		&task.Explanation,
	)

	if err == sql.ErrNoRows {
		return task, nil // Возвращаем пустую запись, если задача не найдена
	} else if err != nil {
		return task, err
	}

	return task, nil
}
