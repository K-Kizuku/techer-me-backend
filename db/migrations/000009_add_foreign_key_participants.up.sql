BEGIN;

ALTER TABLE `participants` ADD CONSTRAINT `fkey_participants_table_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`event_id`);

ALTER TABLE `participants` ADD CONSTRAINT `fkey_participants_table_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`);

COMMIT;
