ALTER TABLE confirmations
    DROP CONSTRAINT fk_confirmations_user_id;
ALTER TABLE tasks
    DROP CONSTRAINT fk_tasks_user_id;

DROP TABLE IF EXISTS confirmations;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS tasks;

DROP TYPE user_status;
DROP TYPE task_status;
