CREATE TABLE IF NOT EXISTS `exchanges` (
  `user_id_1` varchar(255),
  `user_id_2` varchar(255),
  `event_id` varchar(255),
  PRIMARY KEY (`user_id_1`, `user_id_2`, `event_id`)
);