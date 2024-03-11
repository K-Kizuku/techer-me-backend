BEGIN;

ALTER TABLE `user_details` DROP FOREIGN KEY `fkey_user_details_table_user_id`;

ALTER TABLE `event_details` DROP FOREIGN KEY `fkey_event_details_table_event_id`;

ALTER TABLE `event_details` DROP FOREIGN KEY `fkey_event_details_table_user_id`;

ALTER TABLE `timelines` DROP FOREIGN KEY `fkey_timelines_table_event_id`;

ALTER TABLE `exchanges` DROP FOREIGN KEY `fkey_exchanges_table_user_id_1`;

ALTER TABLE `exchanges` DROP FOREIGN KEY `fkey_exchanges_table_user_id_2`;

ALTER TABLE `exchanges` DROP FOREIGN KEY `fkey_exchanges_table_event_id`;

COMMIT;