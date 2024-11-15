package repository

import "context"

func (r repository) AssignTopicToStudent(ctx context.Context, studentId int, topicId int) error {
	query := `
        INSERT INTO topic_student (student_id, topic_id)
        VALUES ($1, $2)
    `

	_, err := r.db.ExecContext(ctx, query, studentId, topicId)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) AssignTaskToTopic(ctx context.Context, topicId int, taskId int) error {
	query := `
        INSERT INTO topic_tasks (task_id, topic_id)
        VALUES ($1, $2)
    `

	_, err := r.db.ExecContext(ctx, query, taskId, topicId)
	if err != nil {
		return err
	}

	return nil
}
