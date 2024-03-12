CREATE TABLE IF NOT EXISTS `event_details` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `event_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `owner_id` varchar(255) NOT NULL,
  `started_at` varchar(255) NOT NULL,
  `finished_at` varchar(255) NOT NULL,
  `message` varchar(255),
  `image_url` varchar(255)
);