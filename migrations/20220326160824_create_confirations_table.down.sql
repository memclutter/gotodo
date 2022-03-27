ALTER TABLE confirmations
    DROP CONSTRAINT fk_confirmation_user_id;
DROP TABLE IF EXISTS confirmations;