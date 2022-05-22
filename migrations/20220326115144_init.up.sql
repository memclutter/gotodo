-- ---------------------------------------------------------------------------------------------------------------------
-- Create custom types
-- ---------------------------------------------------------------------------------------------------------------------
CREATE TYPE task_status AS ENUM ('active', 'archive', 'delete');
CREATE TYPE user_status AS ENUM ('pending', 'active', 'delete');
CREATE TYPE access_role AS ENUM ('member', 'admin');
CREATE TYPE group_status AS ENUM ('active', 'archive', 'delete');
CREATE TYPE project_status AS ENUM ('active', 'archive', 'delete');

-- ---------------------------------------------------------------------------------------------------------------------
-- Create tables
-- ---------------------------------------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS users
(
    id            BIGSERIAL,
    first_name    VARCHAR(500) NOT NULL,
    last_name     VARCHAR(500) NOT NULL,
    email         VARCHAR(500) NOT NULL,
    password_hash VARCHAR(500) NOT NULL,
    status        user_status DEFAULT 'pending',
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

CREATE TABLE IF NOT EXISTS groups
(
    id     BIGSERIAL,
    name   VARCHAR(255) NOT NULL,
    status group_status DEFAULT 'active',
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS projects
(
    id       BIGSERIAL,
    group_id BIGINT       NOT NULL,
    name     VARCHAR(255) NOT NULL,
    status   project_status DEFAULT 'active',
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS access
(
    id         BIGSERIAL,
    user_id    BIGINT,
    group_id   BIGINT      DEFAULT NULL,
    project_id BIGINT      DEFAULT NULL,
    role       access_role DEFAULT 'member',
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS statuses
(
    id         BIGSERIAL,
    project_id BIGINT       NOT NULL,
    name       VARCHAR(128) NOT NULL,
    sort_order SMALLINT DEFAULT 0,
    is_final   BOOLEAN  DEFAULT FALSE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS tasks
(
    id               BIGSERIAL,
    project_id       BIGINT NOT NULL,
    title            TEXT,
    description      TEXT,
    date_created     TIMESTAMP   DEFAULT NOW(),
    status           task_status DEFAULT 'active',
    custom_status_id BIGINT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS participants
(
    id      BIGSERIAL,
    task_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS comments
(
    id           BIGSERIAL,
    task_id      BIGINT NOT NULL,
    user_id      BIGINT NOT NULL,
    date_created TIMESTAMP DEFAULT NOW(),
    title        VARCHAR(255),
    message      TEXT,
    PRIMARY KEY (id)
);

-- ---------------------------------------------------------------------------------------------------------------------
-- Create indexes
-- ---------------------------------------------------------------------------------------------------------------------
CREATE UNIQUE INDEX uix_users_email ON users (email);
CREATE UNIQUE INDEX uix_confirmations_token ON confirmations (token);

-- ---------------------------------------------------------------------------------------------------------------------
-- Create foreign keys
-- ---------------------------------------------------------------------------------------------------------------------
ALTER TABLE confirmations
    ADD CONSTRAINT fk_confirmations_user_id FOREIGN KEY (user_id) REFERENCES users (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE projects
    ADD CONSTRAINT fk_projects_group_id FOREIGN KEY (group_id) REFERENCES groups (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE access
    ADD CONSTRAINT fk_access_user_id FOREIGN KEY (user_id) REFERENCES users (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE access
    ADD CONSTRAINT fk_access_group_id FOREIGN KEY (group_id) REFERENCES groups (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE access
    ADD CONSTRAINT fk_access_project_id FOREIGN KEY (project_id) REFERENCES projects (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE statuses
    ADD CONSTRAINT fk_statuses_project_id FOREIGN KEY (project_id) REFERENCES projects (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE tasks
    ADD CONSTRAINT fk_tasks_project_id FOREIGN KEY (project_id) REFERENCES projects (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE tasks
    ADD CONSTRAINT fk_tasks_custom_status_id FOREIGN KEY (custom_status_id) REFERENCES statuses (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE participants
    ADD CONSTRAINT fk_participants_task_id FOREIGN KEY (task_id) REFERENCES tasks (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE participants
    ADD CONSTRAINT fk_participants_user_id FOREIGN KEY (user_id) REFERENCES users (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE comments
    ADD CONSTRAINT fk_comments_task_id FOREIGN KEY (task_id) REFERENCES tasks (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE comments
    ADD CONSTRAINT fk_comments_user_id FOREIGN KEY (user_id) REFERENCES users (id)
        ON UPDATE CASCADE ON DELETE CASCADE;
