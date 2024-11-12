package service

import (
	"DB_course_paper/server/entity"
	"context"
)

type StudentRepository interface {
	CreateStudent(ctx context.Context, student entity.Student) (int, error)
	DeleteStudent(ctx context.Context, studentId int) error
	UpdateStudent(ctx context.Context, student entity.Student) (int, error)
	GetStudentById(ctx context.Context, studentId int) (entity.Student, error)

	GetStudentTopics(ctx context.Context, studentId int) ([]entity.Topic, error)
}

type TeacherRepository interface {
	CreateTeacher(ctx context.Context, teacher entity.Teacher) (int, error)
	DeleteTeacher(ctx context.Context, teacherId int) error
	UpdateTeacher(ctx context.Context, student entity.Student) (int, error)
	GetStudents(ctx context.Context, teacherId int) ([]entity.Student, error)

	GetTeacherTopics(ctx context.Context, teacherId int) ([]entity.Topic, error)
}

type TokenRepository interface {
	CreateTokens(ctx context.Context, token entity.Token) (int, error)
	DeleteTokens(ctx context.Context, tokenId int) error
	UpdateTokens(ctx context.Context, token entity.Token) (int, error)
	GetTokens(ctx context.Context, tokenId int) ([]entity.Token, error)
}

type GrammarRepository interface {
	CreateGrammarTask(ctx context.Context, task entity.GrammarTask) (int, error)
	DeleteGrammarTask(ctx context.Context, taskId int) error
	GetGrammarTaskById(ctx context.Context, taskId int) (entity.GrammarTask, error)
}

type VocabluaryRepository interface {
	CreateOptionsTask(ctx context.Context, task entity.VocabularyOptionsTask) (int, error)
	DeleteOptionsTask(ctx context.Context, taskId int) error
	GetOptionsTaskById(ctx context.Context, taskId int) (entity.VocabularyOptionsTask, error)

	CreateWordTask(ctx context.Context, task entity.VocabularyWordTask) (int, error)
	DeleteWordTask(ctx context.Context, taskId int) error
	GetWordTaskById(ctx context.Context, taskId int) (entity.VocabularyWordTask, error)
}

type TopicRepository interface {
	CreateTopic(ctx context.Context, topic entity.Topic) (int, error)
	DeleteTopic(ctx context.Context, topicId int) error
	GetTopicById(ctx context.Context, topicId int) (entity.Topic, error)
	UpdateTopic(ctx context.Context, topic entity.Topic) (int, error)
}

type AssignRepository interface {
	AssignTopicToStudent(ctx context.Context, studentId int, topicId int) error
}

type Repositorier interface {
	StudentRepository
}
