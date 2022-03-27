CREATE TABLE IF NOT EXISTS sessions
(
    id           BIGSERIAL,
    token        VARCHAR(64),
    user_id      BIGINT NOT NULL,
    date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    date_expired TIMESTAMP DEFAULT CURRENT_TIMESTAMP + '30 days',
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX uix_sessions_token ON sessions(token);

ALTER TABLE sessions
    ADD CONSTRAINT fk_sessions_user_id FOREIGN KEY (user_id) REFERENCES users (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
