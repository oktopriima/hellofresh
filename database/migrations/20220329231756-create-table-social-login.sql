-- +migrate Up
CREATE TABLE IF NOT EXISTS `hellofresh`.`social_login`
(
    `id`           INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`      INT UNSIGNED NOT NULL,
    `account_type` VARCHAR(45)  NOT NULL,
    `account_id`   VARCHAR(255) NOT NULL,
    `created_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `fk_social_login_users_idx` (`user_id` ASC) VISIBLE,
    CONSTRAINT `fk_social_login_users`
        FOREIGN KEY (`user_id`)
            REFERENCES `hellofresh`.`users` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `hellofresh`.`social_login`;