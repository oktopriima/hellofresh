-- +migrate Up
CREATE TABLE IF NOT EXISTS `hellofresh`.`images`
(
    `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `path`       VARCHAR(255) NOT NULL,
    `mime`       VARCHAR(45)  NOT NULL,
    `file_name`  TEXT         NOT NULL,
    `is_deleted` TINYINT      NOT NULL DEFAULT 0,
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `hellofresh`.`images`;