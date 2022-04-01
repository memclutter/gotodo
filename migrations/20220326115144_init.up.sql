CREATE TABLE IF NOT EXISTS tasks
(
    id           BIGSERIAL,
    user_id      BIGINT NOT NULL,
    title        TEXT,
    description  TEXT,
    date_created TIMESTAMP DEFAULT NOW(),
    status       SMALLINT,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users
(
    id      BIGSERIAL,
    email   VARCHAR(500) NOT NULL,
    pw_hash VARCHAR(500) NOT NULL,
    status  SMALLINT DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS confirmations
(
    id           BIGSERIAL,
    user_id      BIGINT      NOT NULL,
    token        VARCHAR(64) NOT NULL,
    date_created TIMESTAMP DEFAULT current_timestamp,
    date_expired TIMESTAMP DEFAULT current_timestamp + '12 hours',
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX uix_users_email ON users (email);
CREATE UNIQUE INDEX uix_confirmations_token ON confirmations (token);

ALTER TABLE tasks
    ADD CONSTRAINT fk_tasks_user_id FOREIGN KEY (user_id) REFERENCES users (id)
        ON UPDATE CASCADE ON DELETE RESTRICT;
ALTER TABLE confirmations
    ADD CONSTRAINT fk_confirmations_user_id FOREIGN KEY (user_id) REFERENCES users (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
