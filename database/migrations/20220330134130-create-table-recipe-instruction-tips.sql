-- +migrate Up
CREATE TABLE IF NOT EXISTS `hellofresh`.`recipe_instruction_tips`
(
    `id`                    INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `recipe_instruction_id` INT UNSIGNED NOT NULL,
    `value`                 TEXT         NOT NULL,
    `is_deleted`            TINYINT      NOT NULL,
    `created_at`            TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`            TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `fk_recipe_instruction_tips_recipe_instructions1_idx` (`recipe_instruction_id` ASC) VISIBLE,
    CONSTRAINT `fk_recipe_instruction_tips_recipe_instructions1`
        FOREIGN KEY (`recipe_instruction_id`)
            REFERENCES `hellofresh`.`recipe_instructions` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `hellofresh`.`recipe_instruction_tips`;