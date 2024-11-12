package service

import (
	"DB_course_paper/server/entity"
	"context"
)

func (s service) CreateOptionsTask(ctx context.Context, task entity.VocabularyOptionsTask) (int, error) {
	taskId, err := s.r.CreateOptionsTask(ctx, task)
	if err != nil {
		return 0, err
	}
	return taskId, nil
}

func (s service) DeleteOptionsTask(ctx context.Context, taskId int) error {
	err := s.r.DeleteOptionsTask(ctx, taskId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) GetOptionsTaskById(ctx context.Context, taskId int) (entity.VocabularyOptionsTask, error) {
	task, err := s.r.GetOptionsTaskById(ctx, taskId)
	if err != nil {
		return entity.VocabularyOptionsTask{}, err
	}
	return task, nil
}

func (s service) CreateWordTask(ctx context.Context, task entity.VocabularyWordTask) (int, error) {
	taskId, err := s.r.CreateWordTask(ctx, task)
	if err != nil {
		return 0, err
	}
	return taskId, nil
}

func (s service) DeleteWordTask(ctx context.Context, taskId int) error {
	err := s.r.DeleteWordTask(ctx, taskId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) GetWordTaskById(ctx context.Context, taskId int) (entity.VocabularyWordTask, error) {
	task, err := s.r.GetWordTaskById(ctx, taskId)
	if err != nil {
		return entity.VocabularyWordTask{}, err
	}
	return task, nil
}
