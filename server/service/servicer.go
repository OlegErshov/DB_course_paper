package service

import (
	"DB_course_paper/server/entity"
	"context"
)

type StudentService interface {
	CreateStudent(ctx context.Context, student entity.Student) (int, error)
	DeleteStudent(ctx context.Context, studentId int) error
	UpdateStudent(ctx context.Context, student entity.Student) (int, error)
	GetStudentByCreds(ctx context.Context, phone, password string) (int, error)
	LogOutUser(ctx context.Context, userId int) error

	GetStudentTopics(ctx context.Context, studentId int) ([]entity.Topic, error)
}

type TeacherService interface {
	CreateTeacher(ctx context.Context, teacher entity.Teacher) (int, error)
	DeleteTeacher(ctx context.Context, teacherId int) error
	UpdateTeacher(ctx context.Context, student entity.Teacher) (int, error)
	GetTeacherByCreds(ctx context.Context, phone, password string) (int, error)
	GetStudents(ctx context.Context, teacherId int) ([]entity.Student, error)

	GetTeacherTopics(ctx context.Context, teacherId int) ([]entity.Topic, error)
}

type TokenService interface {
	CreateTokens(ctx context.Context, token entity.Token) (int, error)
	DeleteTokens(ctx context.Context, tokenId int) error
	UpdateTokens(ctx context.Context, token entity.Token) (int, error)
	GetTokens(ctx context.Context, tokenId int) (entity.Token, error)

	LogOutUser(ctx context.Context, userId int) error
}

type GrammarService interface {
	CreateGrammarTask(ctx context.Context, task entity.GrammarTask) (int, error)
	DeleteGrammarTask(ctx context.Context, taskId int) error
	GetGrammarTaskById(ctx context.Context, taskId int) (entity.GrammarTask, error)
}

type VocabluaryService interface {
	CreateOptionsTask(ctx context.Context, task entity.VocabularyOptionsTask) (int, error)
	DeleteOptionsTask(ctx context.Context, taskId int) error
	GetOptionsTaskById(ctx context.Context, taskId int) (entity.VocabularyOptionsTask, error)

	CreateWordTask(ctx context.Context, task entity.VocabularyWordTask) (int, error)
	DeleteWordTask(ctx context.Context, taskId int) error
	GetWordTaskById(ctx context.Context, taskId int) (entity.VocabularyWordTask, error)
}

type TopicService interface {
	CreateTopic(ctx context.Context, topic entity.Topic) (int, error)
	DeleteTopic(ctx context.Context, topicId int) error
	GetTopicById(ctx context.Context, topicId int) (entity.Topic, error)
	UpdateTopic(ctx context.Context, topic entity.Topic) (int, error)
}

type AssignService interface {
	AssignTopicToStudent(ctx context.Context, studentId int, topicId int) error
}

type Servicer interface {
	StudentService
	TeacherService
	GrammarService
	VocabluaryService
	TopicService
	AssignService
	TokenService
}
