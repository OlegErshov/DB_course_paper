-- +goose Up
-- +goose StatementBegin
CREATE TABLE role (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30)
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    role_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(20) UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES role(id)
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

CREATE TABLE journal(
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
    hint TEXT
);

CREATE TABLE vocabulary_sentence_task(
    id SERIAL PRIMARY KEY,
    first_part TEXT,
    second_PART TEXT
);

CREATE TABLE vocabulary_options_task(
    id SERIAL PRIMARY KEY,
    sentence TEXT,
    answer_options TEXT
);

CREATE TABLE vocabulary_word_task(
    id SERIAL PRIMARY KEY,
    sentence TEXT,
    answer TEXT
);

CREATE TABLE functional_task(
    id SERIAL PRIMARY KEY,
    sentence TEXT,
    answer TEXT
);

CREATE TABLE topic(
    id SERIAL PRIMARY KEY,
    type TEXT

);

CREATE TABLE topic_tasks(
    id SERIAL PRIMARY KEY,
    task_id INT,
    topic_id INT,
    FOREIGN KEY (topic_id) REFERENCES topic(id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
