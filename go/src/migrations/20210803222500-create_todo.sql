
-- +migrate Up
CREATE TABLE IF NOT EXISTS `todo` (
  `id` varchar(64) NOT NULL,
  `text` varchar(255) NOT NULL,
  `done` tinyint(1) NOT NULL,
  `user_id` varchar(64) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE IF EXISTS `todo`;