CREATE TABLE IF NOT EXISTS `participants` (
  `user_id` varchar(255),
  `event_id` varchar(255),
  PRIMARY KEY (`user_id`, `event_id`)
);