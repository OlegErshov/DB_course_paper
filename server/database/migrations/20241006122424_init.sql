-- +goose Up
-- +goose StatementBegin
CREATE TABLE role (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30)
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE students(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(20) UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE teachers(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(20) UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tokens (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    access_token TEXT,
    refresh_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)

);

CREATE TABLE groups(
    id SERIAL PRIMARY KEY,
    name TEXT,
    students_count INT, --тригер на обновление значения при удалении и добавлении студента в группу
    teacher_id INT NOT NULL,
    FOREIGN KEY (teacher_id) REFERENCES teachers(id)
);

CREATE TABLE results_journal(
    id SERIAL PRIMARY KEY,
    students_id INT,
    topic_id INT UNIQUE
);

CREATE TABLE action_journal(
    id SERIAL PRIMARY KEY,
    action TEXT,
    user_id INT,
    user_role INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE grammar_task(
    id SERIAL PRIMARY KEY,
    sentence TEXT,
    right_answer TEXT,
    hint TEXT,
    explanation TEXT
);

CREATE TABLE vocabulary_sentence_task(
    id SERIAL PRIMARY KEY,
    first_part TEXT,
    second_PART TEXT,
    explanation TEXT
);

CREATE TABLE vocabulary_options_task(
    id SERIAL PRIMARY KEY,
    sentence TEXT,
    answer_options TEXT,
    explanation TEXT
);

CREATE TABLE vocabulary_word_task(
    id SERIAL PRIMARY KEY,
    sentence TEXT,
    answer TEXT,
    explanation TEXT
);

CREATE TABLE functional_task(
    id SERIAL PRIMARY KEY,
    sentence TEXT,
    answer TEXT,
    explanation TEXT
);

CREATE TABLE task(
    id SERIAL PRIMARY KEY,
    type VARCHAR(20),
    exact_task_id  INT UNIQUE,
    FOREIGN KEY (id) REFERENCES results_journal(topic_id)
    -- в эту таблицу пишется только по тригеру и удаляется только по тригеру. Таблица служит агрегатором для всех типов заданий
);

CREATE TABLE topic(
    id SERIAL PRIMARY KEY,
    name TEXT,
    mark INT
);

CREATE TABLE topic_tasks(
    id SERIAL PRIMARY KEY,
    task_id INT,
    topic_id INT,
    FOREIGN KEY (topic_id) REFERENCES topic(id),
    FOREIGN KEY (task_id) REFERENCES task(exact_task_id)
);

CREATE TABLE topic_student(
    id SERIAL PRIMARY KEY,
    student_id INT,
    topic_id INT,
    FOREIGN KEY (topic_id) REFERENCES topic(id),
    FOREIGN KEY (student_id) REFERENCES students(id)
);

CREATE TABLE test(
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE test_topics(
    id SERIAL PRIMARY KEY,
    topic_id INT,
    test_id INT,
    FOREIGN KEY (test_id) REFERENCES test(id),
    FOREIGN KEY (topic_id) REFERENCES topic(id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
