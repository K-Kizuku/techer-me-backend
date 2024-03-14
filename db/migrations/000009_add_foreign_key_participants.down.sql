
BEGIN;

ALTER TABLE `participants` DROP FOREIGN KEY `fkey_participants_table_event_id`;

ALTER TABLE `participants` DROP FOREIGN KEY `fkey_participants_table_user_id`;

COMMIT;