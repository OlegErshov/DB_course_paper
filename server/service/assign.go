package service

import "context"

func (s service) AssignTopicToStudent(ctx context.Context, studentId int, topicId int) error {
	err := s.r.AssignTopicToStudent(ctx, studentId, topicId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) AssignTaskToTopic(ctx context.Context, taskId int, topicId int) error {
	err := s.r.AssignTaskToTopic(ctx, topicId, taskId)
	if err != nil {
		return err
	}
	return nil
}
