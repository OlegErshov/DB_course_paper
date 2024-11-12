package service

import (
	"DB_course_paper/server/entity"
	"context"
)

func (s service) CreateStudent(ctx context.Context, student entity.Student) (int, error) {
	studentId, err := s.r.CreateStudent(ctx, student)
	if err != nil {
		return 0, err
	}
	return studentId, nil
}

func (s service) DeleteStudent(ctx context.Context, studentId int) error {
	err := s.r.DeleteStudent(ctx, studentId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateStudent(ctx context.Context, student entity.Student) (int, error) {
	studentId, err := s.r.UpdateStudent(ctx, student)
	if err != nil {
		return 0, err
	}
	return studentId, nil
}

func (s service) GetStudentByCreds(ctx context.Context, phone, password string) (int, error) {
	student, err := s.r.GetStudentByCreds(ctx, phone, password)
	if err != nil {
		return 0, err
	}
	return student.ID, nil
}

func (s service) GetStudentTopics(ctx context.Context, studentId int) ([]entity.Topic, error) {
	topics, err := s.r.GetStudentTopics(ctx, studentId)
	if err != nil {
		return nil, err
	}
	return topics, nil
}
