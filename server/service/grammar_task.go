package service

import (
	"DB_course_paper/server/entity"
	"context"
)

func (s service) CreateGrammarTask(ctx context.Context, task entity.GrammarTask) (int, error) {
	taskId, err := s.r.CreateGrammarTask(ctx, task)
	if err != nil {
		return 0, err
	}
	return taskId, nil
}

func (s service) DeleteGrammarTask(ctx context.Context, taskId int) error {
	err := s.r.DeleteGrammarTask(ctx, taskId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) GetGrammarTaskById(ctx context.Context, taskId int) (entity.GrammarTask, error) {
	task, err := s.r.GetGrammarTaskById(ctx, taskId)
	if err != nil {
		return entity.GrammarTask{}, err
	}
	return task, nil
}
