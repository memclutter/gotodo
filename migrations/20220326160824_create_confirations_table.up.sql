CREATE TABLE IF NOT EXISTS confirmations
(
    id           BIGSERIAL,
    user_id      BIGINT      NOT NULL,
    token        VARCHAR(64) NOT NULL,
    date_created TIMESTAMP DEFAULT current_timestamp,
    date_expired TIMESTAMP DEFAULT current_timestamp + '12 hours',
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX uix_confirmations_token ON confirmations(token);

ALTER TABLE confirmations
    ADD CONSTRAINT fk_confirmations_user_id FOREIGN KEY (user_id) REFERENCES users (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
