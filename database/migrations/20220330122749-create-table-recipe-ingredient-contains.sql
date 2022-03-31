-- +migrate Up
CREATE TABLE IF NOT EXISTS `hellofresh`.`recipe_ingredient_contains`
(
    `id`                   INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `recipe_ingredient_id` INT UNSIGNED NOT NULL,
    `name`                 VARCHAR(45)  NOT NULL,
    `is_possibility_exist` TINYINT      NOT NULL,
    `is_deleted`           TINYINT      NOT NULL,
    `created_at`           TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`           TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `fk_recipe_ingredient_contains_recipe_ingredients1_idx` (`recipe_ingredient_id` ASC) VISIBLE,
    CONSTRAINT `fk_recipe_ingredient_contains_recipe_ingredients1`
        FOREIGN KEY (`recipe_ingredient_id`)
            REFERENCES `hellofresh`.`recipe_ingredients` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `hellofresh`.`recipe_ingredient_contains`;