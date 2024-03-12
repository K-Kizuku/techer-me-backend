CREATE TABLE IF NOT EXISTS `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` varchar(255) NOT NULL UNIQUE,
  `created_at` varchar(255),
  `updated_at` varchar(255),
  `deleted_at` varchar(255)
);