-- +migrate Up
CREATE TABLE IF NOT EXISTS `hellofresh`.`recipe_ingredients`
(
    `id`                  INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `recipe_id`           INT UNSIGNED NOT NULL,
    `image_id`            INT UNSIGNED NOT NULL,
    `name`                VARCHAR(255) NOT NULL,
    `quantity`            FLOAT        NOT NULL,
    `unit`                VARCHAR(45)  NOT NULL,
    `notes`               TEXT         NOT NULL,
    `is_delivery_include` TINYINT      NOT NULL,
    `is_deleted`          TINYINT      NOT NULL,
    `created_at`          TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`          TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `fk_recipe_ingredients_recipes1_idx` (`recipe_id` ASC) VISIBLE,
    INDEX `fk_recipe_ingredients_images1_idx` (`image_id` ASC) VISIBLE,
    CONSTRAINT `fk_recipe_ingredients_recipes1`
        FOREIGN KEY (`recipe_id`)
            REFERENCES `hellofresh`.`recipes` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION,
    CONSTRAINT `fk_recipe_ingredients_images1`
        FOREIGN KEY (`image_id`)
            REFERENCES `hellofresh`.`images` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `hellofresh`.`recipe_ingredients`;