-- ---------------------------------------------------------------------------------------------------------------------
-- Foreign keys drop
-- ---------------------------------------------------------------------------------------------------------------------
ALTER TABLE task_participants DROP CONSTRAINT fk_task_participants_user_id;
ALTER TABLE task_participants DROP CONSTRAINT fk_task_participants_task_id;
ALTER TABLE tasks DROP CONSTRAINT fk_tasks_custom_status_id;
ALTER TABLE tasks DROP CONSTRAINT fk_tasks_project_id;
ALTER TABLE statuses DROP CONSTRAINT fk_statuses_project_id;
ALTER TABLE access DROP CONSTRAINT fk_access_project_id;
ALTER TABLE access DROP CONSTRAINT fk_access_group_id;
ALTER TABLE access DROP CONSTRAINT fk_access_user_id;
ALTER TABLE projects DROP CONSTRAINT fk_projects_group_id;
ALTER TABLE confirmations DROP CONSTRAINT fk_confirmations_user_id;

-- ---------------------------------------------------------------------------------------------------------------------
-- Table drops
-- ---------------------------------------------------------------------------------------------------------------------
DROP TABLE IF EXISTS task_participants;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS statuses;
DROP TABLE IF EXISTS access;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS confirmations;
DROP TABLE IF EXISTS users;

-- ---------------------------------------------------------------------------------------------------------------------
-- Custom types drops
-- ---------------------------------------------------------------------------------------------------------------------
DROP TYPE project_status;
DROP TYPE group_status;
DROP TYPE access_role;
DROP TYPE user_status;
DROP TYPE task_status;
