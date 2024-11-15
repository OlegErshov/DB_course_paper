package entity

import "time"

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Student struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Teacher struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Token struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Group struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	StudentsCount int       `json:"students_count"`
	TeacherID     int       `json:"teacher_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ResultsJournal struct {
	ID         int `json:"id"`
	StudentsID int `json:"students_id"`
	TopicID    int `json:"topic_id"`
}

type ActionJournal struct {
	ID        int       `json:"id"`
	Action    string    `json:"action"`
	UserID    int       `json:"user_id"`
	UserRole  int       `json:"user_role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GrammarTask struct {
	ID          int    `json:"id"`
	Sentence    string `json:"sentence"`
	RightAnswer string `json:"right_answer"`
	Hint        string `json:"hint"`
	Explanation string `json:"explanation"`
}

type VocabularySentenceTask struct {
	ID          int    `json:"id"`
	FirstPart   string `json:"first_part"`
	SecondPart  string `json:"second_part"`
	Explanation string `json:"explanation"`
}

type VocabularyOptionsTask struct {
	ID            int    `json:"id"`
	Sentence      string `json:"sentence"`
	AnswerOptions string `json:"answer_options"`
	RightAnswer   string `json:"right_answer"`
	Explanation   string `json:"explanation"`
}

type VocabularyWordTask struct {
	ID          int    `json:"id"`
	Sentence    string `json:"sentence"`
	Answer      string `json:"answer"`
	Explanation string `json:"explanation"`
}

type FunctionalTask struct {
	ID          int    `json:"id"`
	Sentence    string `json:"sentence"`
	Answer      string `json:"answer"`
	Explanation string `json:"explanation"`
}

type Task struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	ExactTaskID int    `json:"exact_task_id"`
}

type Topic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Mark int    `json:"mark"`
}

type Test struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
