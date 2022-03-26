CREATE TABLE IF NOT EXISTS tasks
(
    id           BIGSERIAL,
    title        TEXT,
    description  TEXT,
    date_created TIMESTAMP DEFAULT NOW(),
    status       SMALLINT,
    PRIMARY KEY (id)
);
