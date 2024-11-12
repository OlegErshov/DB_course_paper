package service

import (
	"DB_course_paper/server/entity"
	"context"
)

func (s service) CreateTopic(ctx context.Context, topic entity.Topic) (int, error) {
	topicId, err := s.r.CreateTopic(ctx, topic)
	if err != nil {
		return 0, err
	}
	return topicId, nil
}

func (s service) DeleteTopic(ctx context.Context, topicId int) error {
	err := s.r.DeleteTopic(ctx, topicId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) GetTopicById(ctx context.Context, topicId int) (entity.Topic, error) {
	topic, err := s.r.GetTopicById(ctx, topicId)
	if err != nil {
		return entity.Topic{}, err
	}
	return topic, nil
}

func (s service) UpdateTopic(ctx context.Context, topic entity.Topic) (int, error) {
	topicId, err := s.r.UpdateTopic(ctx, topic)
	if err != nil {
		return 0, err
	}
	return topicId, nil
}
