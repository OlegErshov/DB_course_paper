package service

import "context"

func (s service) AssignTopicToStudent(ctx context.Context, studentId int, topicId int) error {
	err := s.r.AssignTopicToStudent(ctx, studentId, topicId)
	if err != nil {
		return err
	}
	return nil
}
