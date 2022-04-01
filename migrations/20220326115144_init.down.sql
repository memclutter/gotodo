ALTER TABLE confirmations
    DROP CONSTRAINT fk_confirmation_user_id;
ALTER TABLE sessions
    DROP CONSTRAINT fk_sessions_user_id;
ALTER TABLE tasks
    DROP CONSTRAINT fk_tasks_user_id;

DROP TABLE IF EXISTS confirmations;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS tasks;
