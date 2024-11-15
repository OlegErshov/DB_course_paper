-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION create_user_from_student()
    RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO users (real_id, role_id,created_at , updated_at)
    VALUES (NEW.ID,1, NOW(), NOW());
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_user_from_student_trigger
    AFTER INSERT ON students
    FOR EACH ROW
EXECUTE FUNCTION create_user_from_student();




CREATE OR REPLACE FUNCTION create_user_from_teacher()
    RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO users (real_id,role_id,created_at, updated_at)
    VALUES (NEW.ID,2, NOW(), NOW());
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_user_from_teacher_trigger
    AFTER INSERT ON teachers
    FOR EACH ROW
EXECUTE FUNCTION create_user_from_teacher();


CREATE OR REPLACE FUNCTION task_aggregation()
    RETURNS TRIGGER AS $$
DECLARE
    type TEXT;
BEGIN
    IF TG_TABLE_NAME = 'vocabulary_options_task' THEN
        type := 'vocabulary_options';
    ELSIF TG_TABLE_NAME = 'vocabulary_word_task' THEN
        type := 'vocabulary_word';
    ELSIF TG_TABLE_NAME = 'grammar_task' THEN
        type := 'grammar';
    ELSE
        RAISE EXCEPTION 'Unsupported table: %', TG_TABLE_NAME;
    END IF;

    INSERT INTO task (exact_task_id, type)
    VALUES (NEW.id, type);

    RETURN NEW;
END
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_grammar_task
    AFTER INSERT ON grammar_task
    FOR EACH ROW
EXECUTE FUNCTION task_aggregation();

CREATE TRIGGER create_vocabulary_options_task
    AFTER INSERT ON vocabulary_options_task
    FOR EACH ROW
EXECUTE FUNCTION task_aggregation();

CREATE TRIGGER create_vocabulary_word_task
    AFTER INSERT ON vocabulary_word_task
    FOR EACH ROW
EXECUTE FUNCTION task_aggregation();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
