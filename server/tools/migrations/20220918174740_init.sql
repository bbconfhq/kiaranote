-- +goose Up
CREATE TABLE IF NOT EXISTS `user`
(
    `id`                  BIGINT    AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `username`            VARCHAR(20) NOT NULL,
    `password`            VARCHAR(255) NOT NULL,
    `last_login_dt`       DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `create_dt`           DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `delete_dt`           DATETIME,

    CONSTRAINT `user_username_uniq` UNIQUE (`username`),

    INDEX `user_username` (`username`)
)
ENGINE = InnoDB
DEFAULT CHARSET = utf8mb4;

-- +goose Down
DROP TABLE user;
