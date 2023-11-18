CREATE TABLE IF NOT EXISTS tasks
(
    id           BIGSERIAL,
    title        TEXT,
    body         TEXT,
    is_done      BOOLEAN   DEFAULT FALSE,
    date_created TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (id)
);
