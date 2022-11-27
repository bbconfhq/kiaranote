-- noinspection SqlNoDataSourceInspectionForFile

-- noinspection SqlDialectInspectionForFile

-- +goose Up
CREATE TABLE IF NOT EXISTS `user`
(
    `id`                  BIGINT        NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `username`            VARCHAR(20)   NOT NULL,
    `password`            VARCHAR(255)  NOT NULL,
    `role`                VARCHAR(20)   NOT NULL DEFAULT 'GUEST',
    `last_login_dt`       DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `create_dt`           DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_dt`           DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `delete_dt`           DATETIME      NULL,

    UNIQUE  (`username`),
    INDEX   (`username`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `audit_log`
(
    `id`                  BIGINT        NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id`             BIGINT        NOT NULL,
    `action_type`         VARCHAR(20)   NOT NULL,
    `action_target`       VARCHAR(255)  NOT NULL,
    `create_dt`           DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`),
    INDEX   (`action_type`),
    INDEX   (`action_target`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `note`
(
    `id`                  BIGINT        NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id`             BIGINT        NOT NULL,
    `title`               VARCHAR(255)  NOT NULL DEFAULT 'Untitled',
    `content`             TEXT          NOT NULL,
    `is_private`          BOOL          NOT NULL DEFAULT FALSE,
    `create_dt`           DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_dt`           DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `delete_dt`           DATETIME      NULL,

    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`),
    INDEX   (`user_id`, `is_private`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `comment`
(
    `id`                  BIGINT        NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id`             BIGINT        NOT NULL,
    `note_id`             BIGINT        NOT NULL,
    `content`             TEXT          NOT NULL,
    `parent_id`           BIGINT        NULL,
    `create_dt`           DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_dt`           DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `delete_dt`           DATETIME      NULL,

    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`),
    FOREIGN KEY (`note_id`) REFERENCES `note`(`id`),
    FOREIGN KEY (`parent_id`) REFERENCES `comment`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `shared_note`
(
    `id`                  BIGINT        NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id`             BIGINT        NOT NULL,
    `note_id`             BIGINT        NOT NULL,
    `password`            VARCHAR(255)  NOT NULL DEFAULT '',
    `create_dt`           DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `expire_dt`           DATETIME      NULL,

    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`),
    FOREIGN KEY (`note_id`) REFERENCES `note`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `note_hierarchy`
(
    `note_id`             BIGINT        NOT NULL,
    `parent_note_id`      BIGINT        NULL,
    `order`               BIGINT        NOT NULL DEFAULT 1,

    FOREIGN KEY (`note_id`) REFERENCES `note`(`id`),
    FOREIGN KEY (`parent_note_id`) REFERENCES `note`(`id`),
    UNIQUE (`note_id`, `parent_note_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- +goose Down
DROP TABLE note_hierarchy;
DROP TABLE shared_note;
DROP TABLE comment;
DROP TABLE note;
DROP TABLE audit_log;
DROP TABLE user;