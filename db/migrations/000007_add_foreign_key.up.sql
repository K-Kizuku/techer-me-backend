BEGIN;

ALTER TABLE `user_details` ADD CONSTRAINT `fkey_user_details_table_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`);

ALTER TABLE `event_details` ADD CONSTRAINT `fkey_event_details_table_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`event_id`);

ALTER TABLE `event_details` ADD CONSTRAINT `fkey_event_details_table_user_id` FOREIGN KEY (`owner_id`) REFERENCES `users` (`user_id`);

ALTER TABLE `timelines` ADD CONSTRAINT `fkey_timelines_table_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`event_id`);

ALTER TABLE `exchanges` ADD CONSTRAINT `fkey_exchanges_table_user_id_1` FOREIGN KEY (`user_id_1`) REFERENCES `users` (`user_id`);

ALTER TABLE `exchanges` ADD CONSTRAINT `fkey_exchanges_table_user_id_2` FOREIGN KEY (`user_id_2`) REFERENCES `users` (`user_id`);

ALTER TABLE `exchanges` ADD CONSTRAINT `fkey_exchanges_table_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`event_id`);

COMMIT;
