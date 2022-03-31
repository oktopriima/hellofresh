-- +migrate Up
CREATE TABLE IF NOT EXISTS `hellofresh`.`recipe_attributes`
(
    `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `recipe_id`  INT UNSIGNED NOT NULL,
    `name`       VARCHAR(255) NOT NULL,
    `color`      VARCHAR(6)   NOT NULL,
    `is_deleted` TINYINT      NOT NULL,
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `fk_recipe_attributes_recipes1_idx` (`recipe_id` ASC) VISIBLE,
    CONSTRAINT `fk_recipe_attributes_recipes1`
        FOREIGN KEY (`recipe_id`)
            REFERENCES `hellofresh`.`recipes` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `hellofresh`.`recipe_attributes`;