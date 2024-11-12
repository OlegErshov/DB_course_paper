package service

import (
	"DB_course_paper/server/entity"
	"context"
)

func (s service) CreateTeacher(ctx context.Context, teacher entity.Teacher) (int, error) {
	teacherId, err := s.r.CreateTeacher(ctx, teacher)
	if err != nil {
		return 0, err
	}
	return teacherId, nil
}

func (s service) DeleteTeacher(ctx context.Context, teacherId int) error {
	err := s.r.DeleteTeacher(ctx, teacherId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateTeacher(ctx context.Context, teacher entity.Teacher) (int, error) {
	teacherId, err := s.r.UpdateTeacher(ctx, teacher)
	if err != nil {
		return 0, err
	}
	return teacherId, nil
}

func (s service) GetStudents(ctx context.Context, teacherId int) ([]entity.Student, error) {
	students, err := s.r.GetStudents(ctx, teacherId)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s service) GetTeacherTopics(ctx context.Context, teacherId int) ([]entity.Topic, error) {
	topics, err := s.r.GetTeacherTopics(ctx, teacherId)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (s service) GetTeacherByCreds(ctx context.Context, phone, password string) (int, error) {
	teacher, err := s.r.GetTeacherByCreds(ctx, phone, password)
	if err != nil {
		return 0, err
	}
	return teacher.ID, nil
}
