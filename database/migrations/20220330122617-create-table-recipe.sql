-- +migrate Up
CREATE TABLE IF NOT EXISTS `hellofresh`.`recipes`
(
    `id`                    INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `image_id`              INT UNSIGNED NOT NULL,
    `title`                 VARCHAR(255) NOT NULL,
    `subtitle`              VARCHAR(255) NOT NULL,
    `slug`                  VARCHAR(255) NOT NULL,
    `description`           TEXT         NOT NULL,
    `preparation_time`      INT          NOT NULL,
    `preparation_time_unit` VARCHAR(45)  NOT NULL,
    `difficulty`            VARCHAR(45)  NOT NULL,
    `is_deleted`            TINYINT      NOT NULL,
    `created_at`            TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`            TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `fk_recipes_images1_idx` (`image_id` ASC) VISIBLE,
    CONSTRAINT `fk_recipes_images1`
        FOREIGN KEY (`image_id`)
            REFERENCES `hellofresh`.`images` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `hellofresh`.`recipes`;