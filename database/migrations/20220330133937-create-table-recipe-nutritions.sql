-- +migrate Up
CREATE TABLE IF NOT EXISTS `hellofresh`.`recipe_nutrition`
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `recipe_id`   INT UNSIGNED NOT NULL,
    `name`        VARCHAR(45)  NOT NULL,
    `per_serving` FLOAT        NOT NULL,
    `per_100g`    FLOAT        NOT NULL,
    `unit`        VARCHAR(45)  NOT NULL,
    `is_deleted`  TINYINT      NOT NULL,
    `created_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `fk_recipe_nutritions_recipes1_idx` (`recipe_id` ASC) VISIBLE,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_recipe_nutritions_recipes1`
        FOREIGN KEY (`recipe_id`)
            REFERENCES `hellofresh`.`recipes` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `hellofresh`.`recipe_nutrition`;