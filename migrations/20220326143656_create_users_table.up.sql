CREATE TABLE IF NOT EXISTS users
(
    id      BIGSERIAL,
    email   VARCHAR(500) NOT NULL,
    pw_hash VARCHAR(500) NOT NULL,
    status  SMALLINT DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX uix_users_email ON users(email);

ALTER TABLE tasks
    ADD COLUMN user_id BIGINT NOT NULL DEFAULT 0;
ALTER TABLE tasks
    ADD CONSTRAINT fk_tasks_user_id FOREIGN KEY (user_id) REFERENCES users (id)
        ON UPDATE CASCADE ON DELETE RESTRICT;
