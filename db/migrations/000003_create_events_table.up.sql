CREATE TABLE IF NOT EXISTS `events` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `event_id` varchar(255) UNIQUE,
  `created_at` varchar(255),
  `updated_at` varchar(255),
  `deleted_at` varchar(255)
);