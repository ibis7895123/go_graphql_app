
-- +migrate Up
ALTER TABLE `todo` ADD FOREIGN KEY `fk_user_id`(`user_id`) REFERENCES `user`(`id`);

-- +migrate Down
ALTER TABLE `todo` DROP FOREIGN KEY `todo_ibfk_1`;