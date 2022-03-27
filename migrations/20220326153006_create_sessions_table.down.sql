ALTER TABLE sessions
    DROP CONSTRAINT fk_sessions_user_id;
DROP TABLE IF EXISTS sessions;