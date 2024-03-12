CREATE TABLE IF NOT EXISTS `user_details` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `is_organizer` bool NOT NULL DEFAULT false,
  `image_url` varchar(255),
  `urls` json NOT NULL,
  `skills` json NOT NULL,
  `message` varchar(255)
);